package terminalport

import "github.com/sashabaranov/go-openai"

type Terminal interface {
	GetUserInput() (string, error)
	PrintReply(message string, rateLimitInsights *openai.RateLimitHeaders)
	Script(script interface{})
	ScriptResults(result interface{})
	Reply(response interface{})
	Error(err error)
}
