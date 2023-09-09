package terminal

import (
	"os"

	"golang.org/x/term"
)

func GetTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	return width, height, err
}
