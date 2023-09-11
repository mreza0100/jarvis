package tools

import (
	"math"
	"os"

	"golang.org/x/term"
)

func GetTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	return width, height, err
}

func calculateTokens(messages []string) int {
	const tokensPer1000Chars = 1333.33

	characters := 0
	for _, m := range messages {
		characters += len(m)
	}

	tokens := int(math.Ceil(float64(characters) / 1000 * tokensPer1000Chars))
	return tokens
}
