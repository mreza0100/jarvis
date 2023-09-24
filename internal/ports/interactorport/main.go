package interactorport

type Interactor interface {
	GetUserInput() (string, error)
	Message(message string, usedTokens int)
	Script(script interface{})
	ScriptResults(result interface{})
	Reply(response interface{})
	Error(err error)
}
