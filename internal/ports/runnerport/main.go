package runnerport

import "github.com/mreza0100/gptjarvis/internal/models"

type Runner interface {
	ExecuteScript(request *models.ScriptRequest) (string, uint8, error)
}
