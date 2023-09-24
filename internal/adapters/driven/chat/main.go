package chat

import (
	"context"
	"encoding/json"
	"math"

	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/pkg/errors"
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

func (c *chat) RawPrompt(rawPrompt string, replyAnswer chatport.Reply, options *chatport.PromptOptions) error {
	ctx := context.Background()

	c.appendMessage(&openai.ChatCompletionMessage{
		Role:    options.PromptRole,
		Content: rawPrompt,
	})

	chat, err := c.clinet.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo16K,
			Messages: c.messages,
		},
	)
	if err != nil {
		return err
	}

	rawReply := chat.Choices[0].Message.Content
	c.appendMessage(&openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: rawReply,
	})

	return json.Unmarshal([]byte(rawReply), replyAnswer)
}

func (c *chat) Prompt(prompt interface{}, replyAnswer chatport.Reply, optionsArg ...*chatport.PromptOptions) error {
	options := c.normalizeOptions(optionsArg...)
	return c.prompt(prompt, replyAnswer, options)
}

func (c *chat) prompt(prompt interface{}, replyAnswer chatport.Reply, options *chatport.PromptOptions) error {
	rawPrompt, err := json.Marshal(prompt)
	if err != nil {
		return errors.Wrap(err, "failed to parse jarvis json reply, Jarvis must provide a valid json")
	}
	return c.RawPrompt(string(rawPrompt), replyAnswer, options)
}

func (c *chat) appendMessage(chat *openai.ChatCompletionMessage) {
	c.messages = append(c.messages, *chat)
}

func (c *chat) CountTokens() int {
	const tokensPer1000Chars = 1333.33

	characters := 0
	for _, m := range c.messages {
		characters += len(m.Content)
	}

	tokens := int(math.Ceil(float64(characters) / 1000 * tokensPer1000Chars))
	return tokens
}
