package chatport

type PromptOptions struct {
	PromptRole string
}

// TODO: Change this name to GPTServer or something else then Chat!
type Chat interface {
	Prompt(prompt interface{}, replyAnswer interface{}, optionsArg ...*PromptOptions) error
	RawPrompt(rawPrompt string, replyAnswer interface{}, options *PromptOptions) error
}
