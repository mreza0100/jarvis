package chatport

import "github.com/mreza0100/gptjarvis/internal/models"

type PromptOptions struct {
	PromptRole string
}

type Chat interface {
	Prompt(prompt *models.Prompt, optionsArg ...*PromptOptions) (*models.Response, error)
}
