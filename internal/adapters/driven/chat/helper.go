package chat

import (
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/sashabaranov/go-openai"
)

func (c *chat) normalizeOptions(optionsArg ...*chatport.PromptOptions) *chatport.PromptOptions {
	defaultOptions := &chatport.PromptOptions{
		PromptRole: openai.ChatMessageRoleSystem,
	}

	var options *chatport.PromptOptions
	if len(optionsArg) != 0 {
		options = c.normalizeOptions(optionsArg[0])
	} else {
		return defaultOptions
	}

	if options.PromptRole != "" {
		defaultOptions.PromptRole = options.PromptRole
	}

	return options
}
