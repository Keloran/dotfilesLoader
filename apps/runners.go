package apps

import (
	"fmt"
	"os/exec"

	"github.com/Keloran/dotfilesLoader/console"
)

type Runner interface {
	install() error
}

type Brew struct {
	App     string
	Console console.Console
}

type Cask struct {
	App     string
	Console console.Console
}

type MAS struct {
	App     string
	Console console.Console
}

type APM struct {
	Package string
	Console console.Console
}

func Install(r Runner) error {
	return r.install()
}

func (i Brew) install() error {
	out, err := exec.Command("brew", "install", i.App).Output()
	if err != nil {
		return fmt.Errorf("runners brew: %w", err)
	}

	if len(out) > 0 {
		i.Console.Log(fmt.Sprintf("%s output: %s\n", i.App, out))
	} else {
		i.Console.Info(fmt.Sprintf("%s installed\n", i.App))
	}

	return nil
}

func (i Cask) install() error {
	out, err := exec.Command("brew", "cask", "install", i.App).Output()
	if err != nil {
		return fmt.Errorf("runners casks: %w", err)
	}

	if len(out) > 0 {
		i.Console.Log(fmt.Sprintf("%s output: %s\n", i.App, out))
	} else {
		i.Console.Info(fmt.Sprintf("%s installed\n", i.App))
	}

	return nil
}

func (i MAS) install() error {
	out, err := exec.Command("mas", "install", i.App).Output()
	if err != nil {
		return fmt.Errorf("runners mas: %w", err)
	}

	if len(out) > 0 {
		i.Console.Log(fmt.Sprintf("%s output: %s\n", i.App, out))
	} else {
		i.Console.Info(fmt.Sprintf("%s installed\n", i.App))
	}

	return nil
}

func (i APM) install() error {
	out, err := exec.Command("apm", "install", i.Package).Output()
	if err != nil {
		return fmt.Errorf("runners mas: %w", err)
	}

	if len(out) > 0 {
		i.Console.Log(fmt.Sprintf("%s output: %s\n", i.Package, out))
	} else {
		i.Console.Info(fmt.Sprintf("%s installed\n", i.Package))
	}

	return nil
}
