package interactorport

import "github.com/mreza0100/gptjarvis/internal/models"

type Interactor interface {
	GetUserInput() (string, error)
	Message(message string)
	Script(script *models.ScriptRequest)
	ScriptResults(result *models.ScriptResult)
	Response(response *models.Response)
}
