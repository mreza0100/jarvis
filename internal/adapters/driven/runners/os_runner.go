package runners

import (
	"errors"
	"os/exec"
	"strings"
	"syscall"

	"github.com/mreza0100/jarvis/internal/models"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
)

type osRunner struct{}

func NewOSRunner() runnerport.OSRunner {
	return &osRunner{}
}

func (b *osRunner) ExecScript(req *models.OSRunnerRequest) (*models.OSRunnerResult, error) {
	cmd := exec.Command(req.Runtime)
	cmd.Stdin = strings.NewReader(req.Script)

	// output = stdout + stderr
	rawOutput, err := cmd.CombinedOutput()
	output := string(rawOutput)
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				return &models.OSRunnerResult{
					StatusCode: status.ExitStatus(),
					Stdout:     output,
					Stderr:     err.Error(),
				}, nil
			}
		}
		return &models.OSRunnerResult{
			StatusCode: -1,
			Stdout:     output,
			Stderr:     err.Error(),
		}, err
	}

	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		return &models.OSRunnerResult{
			StatusCode: status.ExitStatus(),
			Stdout:     output,
		}, nil
	}

	return &models.OSRunnerResult{
		StatusCode: 0,
		Stdout:     output,
	}, nil
}
