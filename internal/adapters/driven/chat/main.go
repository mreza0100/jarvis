package chat

import (
	"context"
	"encoding/json"

	"github.com/mreza0100/gptjarvis/internal/models"
	"github.com/mreza0100/gptjarvis/internal/pkg/terminal"
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

func (c *chat) rawPrompt(rawPrompt string, options *chatport.PromptOptions) (*models.Response, error) {
	ctx := context.Background()

	c.appendMessage(&openai.ChatCompletionMessage{
		Role:    options.PromptRole,
		Content: rawPrompt,
	})

	resp, err := c.clinet.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo16K,
			Messages: c.messages,
		},
	)
	if err != nil {
		return nil, err
	}

	rawResponse := resp.Choices[0].Message.Content
	c.appendMessage(&openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: rawResponse,
	})

	response := new(models.Response)
	err = json.Unmarshal([]byte(rawResponse), response)
	return response, err
}

func (c *chat) Prompt(prompt *models.Prompt, optionsArg ...*chatport.PromptOptions) (*models.Response, error) {
	options := c.normalizeOptions(optionsArg...)
	return c.prompt(prompt, options)
}

func (c *chat) appendMessage(chat *openai.ChatCompletionMessage) {
	c.messages = append(c.messages, *chat)
}

func (c *chat) prompt(prompt *models.Prompt, options *chatport.PromptOptions) (*models.Response, error) {
	rawPrompt, err := json.Marshal(prompt)
	if err != nil {
		return nil, err
	}

	prompt.Screen, err = terminal.GetTerminalSize()
	if err != nil {
		return nil, err
	}

	response, err := c.rawPrompt(string(rawPrompt), options)
	return response, err
}
