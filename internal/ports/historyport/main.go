package historyport

type History interface {
	SavePrompt(prompt interface{})
	SaveReply(prompt interface{})
}
