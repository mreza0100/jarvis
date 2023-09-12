package runnerport

import "github.com/mreza0100/jarvis/internal/models"

type OSRunner interface {
	ExecScript(req *models.OSRunnerRequest) (*models.OSRunnerResult, error)
}
