package historyport

type History interface {
	SavePrompt(prompt interface{})
	SaveResponse(prompt interface{})
}
