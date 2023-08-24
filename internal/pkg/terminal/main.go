package terminal

import (
	"os"

	"github.com/mreza0100/jarvis/internal/models"
	"golang.org/x/term"
)

func GetTerminalSize() (models.Screen, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))

	return models.Screen{
		Width:  width,
		Height: height,
	}, err
}
