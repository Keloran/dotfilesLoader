package main

import (
	"fmt"

	"github.com/Keloran/dotfilesLoader/console"
)

func Help() {
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
