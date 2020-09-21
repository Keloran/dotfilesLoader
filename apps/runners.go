package apps

import (
	"fmt"
	"os/exec"
)

type Runner interface {
	install() error
}

type Brew struct {
	App string
}

type Cask struct {
	App string
}

type MAS struct {
	App string
}

type APM struct {
	Package string
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
		fmt.Printf("%s\n", out)
	} else {
		fmt.Printf("%s installed\n", i.App)
	}

	return nil
}

func (i Cask) install() error {
	return nil
}

func (i MAS) install() error {
	return nil
}

func (i APM) install() error {
	return nil
}
