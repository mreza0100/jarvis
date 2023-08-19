package chat

import (
	"context"

	"github.com/mreza0100/gptjarvis/internal/ports/chatport"
	"github.com/sashabaranov/go-openai"
)

type chat struct {
	clinet   openai.Client
	messages []openai.ChatCompletionMessage
}

type NewChatReq struct {
	Clinet *openai.Client
}

func NewChat(req *NewChatReq) chatport.Chat {
	return &chat{
		clinet:   *req.Clinet,
		messages: make([]openai.ChatCompletionMessage, 0, 5),
	}
}

func (c *chat) Prompt(prompt string, optionsArg ...*chatport.PromptOptions) (string, error) {
	var options *chatport.PromptOptions
	if len(optionsArg) != 0 {
		options = c.normalizeOptions(optionsArg[0])
	} else {
		options = c.normalizeOptions(nil)
	}

	return c.prompt(prompt, options)
}

func (c *chat) normalizeOptions(options *chatport.PromptOptions) *chatport.PromptOptions {
	defaultOptions := &chatport.PromptOptions{
		PromptRole: openai.ChatMessageRoleUser,
	}
	if options == nil {
		return defaultOptions
	}

	if options.PromptRole != "" {
		defaultOptions.PromptRole = options.PromptRole
	}

	return defaultOptions
}

func (c *chat) prompt(prompt string, options *chatport.PromptOptions) (string, error) {
	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    options.PromptRole,
		Content: prompt,
	})
	ctx := context.Background()

	resp, err := c.clinet.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo16K,
			Messages: c.messages,
		},
	)
	if err != nil {
		return "", err
	}

	answer := resp.Choices[0].Message.Content

	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: answer,
	})

	return answer, err
}
