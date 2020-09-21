package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Keloran/dotfilesLoader/apps"
	"github.com/Keloran/dotfilesLoader/console"
	"github.com/Keloran/dotfilesLoader/dots"
	"github.com/Keloran/dotfilesLoader/files"
)

func main() {
	currentLocation, err := os.Getwd()
	if err != nil {
		fmt.Printf("couldn't get current directory: %+v", err)
		return
	}

	command := ""
	if len(os.Args) >= 2 {
		for _, arg := range os.Args {
			if strings.Contains(arg, "-") {
				continue
			}
			command = arg
		}
	}

	var dotfiles, appsGUI, appsCLI, OS, all bool
	help := true
	all = true

	switch command {
	case "dots":
		fallthrough
	case "dotfiles":
		dotfiles = true
		all = false
		help = false
	case "gui":
		appsGUI = true
		all = false
		help = false
	case "cli":
		appsCLI = true
		help = false
		all = false
	case "help":
		all = false
	}

	// vars
	var USERNAME = ""
	var Github = files.Github{}

	// flags
	force := flag.Bool("force", false, "Force override")
	skip := flag.Bool("skip-download", false, "Skip Download")

	flag.StringVar(&USERNAME, "user", "", "Username to inject")

	flag.StringVar(&Github.Username, "github-user", "keloran", "github username")
	flag.StringVar(&Github.Username, "github-username", "keloran", "github username")

	flag.StringVar(&Github.Repository, "github-repository", "dotfiles", "github repository")
	flag.StringVar(&Github.Repository, "github-repo", "dotfiles", "github repository")
	flag.StringVar(&Github.GivenLocation, "github-files", "testbed", "force location of the files")

	flag.Parse()

	// dots
	dot := dots.Dots{
		Username:   USERNAME,
		Force:      *force,
		Skip:       *skip,
		Github:     Github,
		CurrentDir: currentLocation,
	}
	if dotfiles {
		if err := dot.Install(); err != nil {
			console.NewConsole(false).Error(fmt.Sprintf("dotfiles err: %+v", err))
		}
		return
	}

	// apps
	app := apps.Apps{
		Username: USERNAME,
		Force:    *force,
		Skip:     *skip,
		Github:   Github,
	}
	// GUI
	if appsGUI {
		// need sudo password
		if err := app.SetSudo(); err != nil {
			console.NewConsole(false).Error(fmt.Sprintf("Sudo err: %+v", err))
		}

		// GUI install
		fmt.Print("Apps GUI\n")
		if err = app.InstallGUI(); err != nil {
			console.NewConsole(false).Error(fmt.Sprintf("GUI err: %+v", err))
		}

		fmt.Print("Apps GUI\n")
		return
	}

	// CLI
	if appsCLI {
		// need sudo password
		if err := app.SetSudo(); err != nil {
			console.NewConsole(false).Error(fmt.Sprintf("Sudo err: %+v", err))
		}

		// CLI install
		fmt.Print("Apps CLI\n")
		if err = app.InstallCLI(); err != nil {
			console.NewConsole(false).Error(fmt.Sprintf("CLI err: %v", err))
		}
		return
	}

	// OS
	if OS {
		fmt.Print("OS\n")
		return
	}

	// All
	if all {
		fmt.Print("All\n")
		return
	}

	// Help
	if help {
		Help()
	}
}
