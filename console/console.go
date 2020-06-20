package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start(text string) {
	fmt.Print(Green(fmt.Sprintf("--- Starting %s ---\n", text)))
}

func End(text string) {
	fmt.Print(Green(fmt.Sprintf("*** Ending %s ***\n", text)))
}

func Warning(text string) {
	fmt.Print(Blue(fmt.Sprintf("+++ %s +++\n", text)))
}

func Info(text string) {
	fmt.Print(Cyan(fmt.Sprintf("%s\n", text)))
}

func Error(text string) {
	fmt.Print(Red(fmt.Sprintf("### %s ###\n", text)))
}

func Log(text string) {
	fmt.Print(White(fmt.Sprintf("%s\n", text)))
}

func Nice(text string) {
	fmt.Print(Yellow(fmt.Sprintf("%s\n", text)))
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
		return "", fmt.Errorf("getResponse %s, err: %w", err)
	}
	in = strings.Replace(in, "\n", "", -1)
	if lower {
		in = strings.ToLower(in)
	}
	return in, err
}

func BlankLine() {
	fmt.Printf("\n")
}

func Skipping(text string) {
	Nice(fmt.Sprintf("   Skipping %s", text))
}

func Installing(text string) {
	Nice(fmt.Sprintf(">>> Installing %s ---", text))
}

func Installed(text string) {
	Info(fmt.Sprintf(" %s Installed <<< ", text))
}
