package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/keloran/dotfilesLoader/console"
	"github.com/keloran/dotfilesLoader/dots"
	"github.com/keloran/dotfilesLoader/files"
)

func main() {
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

	flag.Parse()

	// dots
	dot := dots.Dots{
		Username: USERNAME,
		Force:    *force,
		Skip:     *skip,
		Github:   Github,
	}
	if dotfiles {
		err := dot.Install()
		if err != nil {
			c := console.NewConsole(false)
			c.Error(fmt.Sprintf("dotfiles err: %+v", err))
		}
		return
	}

	// GUI
	if appsGUI {
		return
	}

	// CLI
	if appsCLI {
		return
	}

	// OS
	if OS {
		return
	}

	// All
	if all {
		return
	}

	// Help
	if help {
		hFormat := "    %s        %s"
		fFormat := "    -%s        %s"

		c := console.NewConsole(true)
		c.Start("Help").
			Info("Usage dotfiles <command> <flags>").
			BlankLine().
			Info("Commands:").
			FormattedText(fmt.Sprintf(hFormat, console.Cyan("help"), console.Blue("  This message"))).
			FormattedText(fmt.Sprintf(hFormat, console.Cyan("cli"), console.Blue("   Install CLI Apps"))).
			FormattedText(fmt.Sprintf(hFormat, console.Cyan("gui"), console.Blue("   Install GUI Apps"))).
			FormattedText(fmt.Sprintf(hFormat, console.Cyan("os"), console.Blue("    Install OS Settings"))).
			FormattedText(fmt.Sprintf(hFormat, console.Cyan("dots"), console.Blue("  Install the dotfiles"))).
			FormattedText(fmt.Sprintf(hFormat, console.Cyan("update"), console.Blue("Run updaters"))).
			BlankLine().
			Info("Flags:").
			FormattedText(fmt.Sprintf(fFormat, console.Green("github-user | github-username <keloran>"), console.Blue("   Github username"))).
			FormattedText(fmt.Sprintf(fFormat, console.Green("github-repository | github-repo <dotfiles>"), console.Blue("Github repository"))).
			FormattedText(fmt.Sprintf(fFormat, console.Green("force"), console.Blue("                                     Force installs"))).
			FormattedText(fmt.Sprintf(fFormat, console.Green("skip-download"), console.Blue("                             Skip Download"))).
			FormattedText(fmt.Sprintf(fFormat, console.Green("user <keloran>"), console.Blue("                            The username if needs sudo"))).
			End("")
	}
}
