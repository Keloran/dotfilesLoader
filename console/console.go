package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Console struct {
	Chain      bool
	FullText   string
	StartLabel string
}

func NewConsole(chain bool) Console {
	return Console{
		Chain: chain,
	}
}

func (c Console) Start(text string) Console {
	cont := Green(fmt.Sprintf("--- Starting %s ---\n", text))
	c.StartLabel = text

	if c.Chain {
		c.FullText = fmt.Sprint(cont)
	} else {
		fmt.Print(cont)
	}

	return c
}

func (c Console) End(text string) Console {
	if len(text) == 0 {
		text = c.StartLabel
	}

	cont := Green(fmt.Sprintf("\n*** Ending %s ***\n", text))

	if c.Chain {
		c.FullText = fmt.Sprintf("%s %s", c.FullText, cont)
		fmt.Print(c.FullText)
		c.FullText = ""
	} else {
		fmt.Print(cont)
	}

	return c
}

func (c Console) Warning(text string) Console {
	cont := Blue(fmt.Sprintf("+++ %s +++\n", text))

	if c.Chain {
		c.FullText = fmt.Sprintf("%s %s", c.FullText, cont)
	} else {
		fmt.Print(cont)
	}

	return c
}

func (c Console) Info(text string) Console {
	cont := Cyan(fmt.Sprintf("%s\n", text))

	if c.Chain {
		c.FullText = fmt.Sprintf("%s %s", c.FullText, cont)
	} else {
		fmt.Print(cont)
	}

	return c
}

func (c Console) Error(text string) Console {
	cont := Red(fmt.Sprintf("### %s ###\n", text))

	if c.Chain {
		c.FullText = fmt.Sprintf("%s %s", c.FullText, cont)
	} else {
		fmt.Print(cont)
	}

	return c
}

func (c Console) Log(text string) Console {
	cont := White(fmt.Sprintf("%s\n", text))

	if c.Chain {
		c.FullText = fmt.Sprintf("%s %s", c.FullText, cont)
	} else {
		fmt.Print(cont)
	}

	return c
}

func (c Console) Nice(text string) Console {
	cont := Yellow(fmt.Sprintf("%s\n", text))

	if c.Chain {
		c.FullText = fmt.Sprintf("%s %s", c.FullText, cont)
	} else {
		fmt.Print(cont)
	}

	return c
}

func Question(text string, lower bool) (string, error) {
	fmt.Print(Magenta(fmt.Sprintf("%s ?\n", text)))
	fmt.Print(White("~> "))

	return getResponse(lower)
}

func getResponse(lower bool) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("getResponse %s, err: %w", in, err)
	}
	in = strings.Replace(in, "\n", "", -1)
	if lower {
		in = strings.ToLower(in)
	}
	return in, err
}

func (c Console) BlankLine() Console {
	if c.Chain {
		c.FullText = fmt.Sprintf("%s\n", c.FullText)
	} else {
		fmt.Print("\n")
	}

	return c
}

func (c Console) Skipping(text string) Console {
	c.Nice(fmt.Sprintf("   Skipping %s", text))

	return c
}

func (c Console) Installing(text string) Console {
	c.Nice(fmt.Sprintf(">>> Installing %s ---", text))

	return c
}

func (c Console) Installed(text string) Console {
	c.Info(fmt.Sprintf(" %s Installed <<< ", text))

	return c
}

func (c Console) FormattedText(text string) Console {
	if c.Chain {
		c.FullText = fmt.Sprintf("%s %s\n", c.FullText, text)
	} else {
		fmt.Printf("%s %s\n", c.FullText, text)
	}

	return c
}
