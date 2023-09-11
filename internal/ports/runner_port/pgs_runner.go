package runnerport

import "github.com/mreza0100/jarvis/internal/models"

type PgsRunner interface {
	ExecScript(req *models.PgsRunnerRequest) (*models.PgsRunnerResponse, error)
}
