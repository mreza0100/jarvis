package chatport

type PromptOptions struct {
	PromptRole string
}

type Chat interface {
	Prompt(prompt interface{}, replyAnswer interface{}, optionsArg ...*PromptOptions) error
	RawPrompt(rawPrompt string, replyAnswer interface{}, options *PromptOptions) error
	CountTokens() int
}
