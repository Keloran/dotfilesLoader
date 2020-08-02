package console

import (
	"fmt"
)

func Blue(text string) string {
	return console(text, 34)
}

func Green(text string) string {
	return console(text, 32)
}

func Red(text string) string {
	return console(text, 31)
}

func Black(text string) string {
	return console(text, 30)
}

func Yellow(text string) string {
	return console(text, 33)
}

func Magenta(text string) string {
	return console(text, 35)
}

func Cyan(text string) string {
	return console(text, 36)
}

func White(text string) string {
	return console(text, 37)
}

func console(text string, color int) string {
	return fmt.Sprintf("\u001b[%dm%s%s", color, text, reset())
}

func reset() string {
	return fmt.Sprint("\u001b[0m")
}
