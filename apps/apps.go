package apps

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
  "log"
  "os"
	"os/exec"
	"strings"

	"github.com/Keloran/dotfilesLoader/console"
	"github.com/Keloran/dotfilesLoader/files"

	"gopkg.in/yaml.v3"
)

type Apps struct {
	Username string
	files.Github
	Force        bool
	Skip         bool
	SudoPassword string
}

type AppsFile []struct {
	Installer string   `yaml:"installer"`
	Extra     string   `yaml:"installer-extra,omitempty"`
	ToInstall []string `yaml:"apps"`
}

func (a Apps) SetSudo() error {
	c := console.NewConsole(false)
	c.Info("Need your SUDO password")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadBytes(byte(10))
	if err != nil {
		return fmt.Errorf("SetSudo readbytes: %w", err)
	}

	a.SudoPassword = string(input)
	return nil
}

func xcode() error {
	// make sure xcode is installed
	_, err := exec.Command("which", "xcode-select").Output()
	if err != nil {
		return fmt.Errorf("xcode which: %w", err)
	}

	// run xcode select installer
	c := exec.Command("xcode-select", "--install")
	stderr, err := c.StderrPipe()
	if err != nil {
		return fmt.Errorf("xcode select stderr: %w", err)
	}
	if err := c.Start(); err != nil {
		return fmt.Errorf("xcode select start: %w", err)
	}
	s, _ := ioutil.ReadAll(stderr)
	if strings.Contains(string(s), "already installed") {
		// xcode already installed
		return nil
	}
	if err := c.Wait(); err != nil {
		return fmt.Errorf("xcode select wait: %w", err)
	}

	return nil
}

func (a Apps) brewInstall() error {
	// make sure brew is installed
	installed, err := exec.Command("which", "brew").Output()
	if err != nil {
		return fmt.Errorf("brewInstall which: %w", err)
	}

	// not installed
	if !strings.Contains(string(installed), "bin/brew") {
		c := exec.Command("/bin/bash", "-c", `"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"`)
		stdin, err := c.StdinPipe()
		if err != nil {
			return fmt.Errorf("brewInstall stdin: %w", err)
		}

		go func() {
			defer stdin.Close()
			_, err := io.WriteString(stdin, a.SudoPassword)
			if err != nil {
			  log.Fatalf("brew install write sudo: %+v", err)
      }
		}()

		_, err = c.CombinedOutput()
		if err != nil {
			return fmt.Errorf("brewInstall combined: %w", err)
		}
	}

	return nil
}

func (a Apps) InstallCLI() error {
	if err := xcode(); err != nil {
		return fmt.Errorf("cli xcode install: %w", err)
	}

	if err := a.brewInstall(); err != nil {
		return fmt.Errorf("cli brewInstall: %w", err)
	}

	if a.Github.Repository != "" {
		filesLocation, err := files.Downloader{
			GithubDetails: a.Github,
			Skip:          a.Skip,
		}.Github()
		if err != nil {
			return fmt.Errorf("dotfiles install: %w", err)
		}

		if filesLocation == "" {
			return fmt.Errorf("dotfiles install, file location blank")
		}

		dat, err := ioutil.ReadFile(fmt.Sprintf("%s/apps.yml", filesLocation))
		if err != nil {
			return fmt.Errorf("brew readfile: %w", err)
		}

		apps := AppsFile{}
		if err = yaml.Unmarshal(dat, &apps); err != nil {
			return fmt.Errorf("brew yaml unmarshal: %w", err)
		}

		for _, installer := range apps {
			switch installer.Installer {
			case "brew":
				for _, app := range installer.ToInstall {
					if installer.Extra == "tap" {
						if err := tap(app); err != nil {
							return fmt.Errorf("tap install: %w", err)
						}
						continue
					}

					i := Brew{
						App: app,
					}
					if err := Install(i); err != nil {
						return fmt.Errorf("brew install: %w", err)
					}
				}
			case "cask":
				for _, app := range installer.ToInstall {
					i := Cask{
						App: app,
					}
					if err := Install(i); err != nil {
						return fmt.Errorf("cask install: %w", err)
					}
				}
			case "apm":
				for _, app := range installer.ToInstall {
					i := APM{
						Package: app,
					}
					if err := Install(i); err != nil {
						return fmt.Errorf("APM install: %w", err)
					}
				}
			case "mas":
				for _, app := range installer.ToInstall {
					i := MAS{
						App: app,
					}
					if err := Install(i); err != nil {
						return fmt.Errorf("MAS install: %w", err)
					}
				}
			}
		}
	}

	return nil
}

func tap(t string) error {
	out, err := exec.Command("brew", "tap", t).Output()
	if err != nil {
		return fmt.Errorf("apps tap: %w", err)
	}

	if len(out) > 0 {
		fmt.Printf("brew tap output: %s\n", out)
	}

	return nil
}
