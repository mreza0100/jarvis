package models

import (
	"log"

	"github.com/mreza0100/jarvis/internal/pkg/tools"
)

type Screen struct {
	Width  int `json:"Width"`
	Height int `json:"Height"`
}

func (s *Screen) GetScreen() *Screen {
	width, height, err := tools.GetTerminalSize()
	if err != nil {
		log.Fatal(err)
	}

	if s == nil {
		s = &Screen{
			Width:  width,
			Height: height,
		}
		return s
	}

	if isChanged := width != s.Width || height != s.Height; isChanged {
		s.Width, s.Height = width, height
		return s
	}
	return nil
}
