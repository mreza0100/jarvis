package chat

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mreza0100/jarvis/internal/models"
	errs "github.com/mreza0100/jarvis/internal/pkg/errors"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

type chat struct {
	config   *models.ChatConfig
	clinet   *openai.Client
	messages []openai.ChatCompletionMessage
	headers  openai.RateLimitHeaders
}

type NewChatReq struct {
	ChatConfigs *models.ChatConfig
}

func NewChat(req *NewChatReq) chatport.Chat {
	ch := &chat{
		config:   req.ChatConfigs,
		clinet:   openai.NewClient(req.ChatConfigs.GetToken()),
		messages: make([]openai.ChatCompletionMessage, 0, 5),
		headers:  openai.RateLimitHeaders{},
	}

	return ch
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
			Messages: c.messages,

			Model:       c.config.Model,
			Temperature: c.config.Temperature,
		},
	)
	if err != nil {
		e := &openai.APIError{}
		if errs.As(err, &e) {
			switch e.HTTPStatusCode {
			case 401:
				return errs.OpenAPIInvalidAuth(e.Message)
			case 429:
				return errs.OpenAPIRateLimit(e.Message, time.Second*5)
			case 500:
				return errs.OpenAPIInternalError(e.Message)
			default:
				return err
			}
		}
	}

	c.headers = chat.GetRateLimitHeaders()
	rawReply := chat.Choices[0].Message.Content
	c.appendMessage(&openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: rawReply,
	})

	return json.Unmarshal([]byte(rawReply), replyAnswer)
}

func (c *chat) Prompt(prompt chatport.Prompt, replyAnswer chatport.Reply, optionsArg ...*chatport.PromptOptions) error {
	options := c.normalizeOptions(optionsArg...)
	return c.prompt(prompt, replyAnswer, options)
}

func (c *chat) prompt(prompt chatport.Prompt, replyAnswer chatport.Reply, options *chatport.PromptOptions) error {
	rawPrompt, err := json.Marshal(prompt)
	if err != nil {
		return errors.Wrap(err, "failed to parse jarvis json reply, Jarvis must provide a valid json")
	}
	return c.RawPrompt(string(rawPrompt), replyAnswer, options)
}

func (c *chat) appendMessage(chat *openai.ChatCompletionMessage) {
	c.messages = append(c.messages, *chat)
}

func (c *chat) GetRateLimitInsights() *openai.RateLimitHeaders {
	return &c.headers
}
