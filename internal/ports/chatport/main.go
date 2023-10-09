package chatport

type PromptOptions struct {
	PromptRole string
}

type (
	Reply  interface{}
	Prompt interface{}
)

type Chat interface {
	Prompt(prompt Prompt, replyAnswer Reply, optionsArg ...*PromptOptions) error
	RawPrompt(rawPrompt string, replyAnswer Reply, options *PromptOptions) error
	CountTokens() int
}
