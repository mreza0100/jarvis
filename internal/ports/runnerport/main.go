package runnerport

import "github.com/mreza0100/jarvis/internal/models"

type Runner interface {
	ExecuteScript(request *models.ScriptRequest) (string, uint8, error)
}
