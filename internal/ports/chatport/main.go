package chatport

type PromptOptions struct {
	PromptRole string
}

type Chat interface {
	Prompt(prompt string, optionsArg ...*PromptOptions) (string, error)
}
