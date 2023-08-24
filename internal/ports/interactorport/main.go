package interactorport

import "github.com/mreza0100/jarvis/internal/models"

type Interactor interface {
	GetUserInput() (string, error)
	Message(message string, usedTokens int)
	Script(script *models.ScriptRequest)
	ScriptResults(result *models.ScriptResult)
	Response(response *models.Response)
}
