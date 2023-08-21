package runner

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"

	"github.com/gofrs/uuid"
	"github.com/mreza0100/gptjarvis/internal/models"
	"github.com/mreza0100/gptjarvis/internal/ports/runnerport"
)

type runner struct {
	conversationID string
	rootDir        string
}

func NewRunner() runnerport.Runner {
	conversationID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conversationID.String())
	return &runner{
		conversationID: conversationID.String(),
		rootDir:        "/home/mamzi/works/gpt-jarvis/tmp",
	}
}

func (r *runner) ExecuteScript(request *models.ScriptRequest) (string, uint8, error) {
	cmd := exec.Command(request.Runtime)
	cmd.Stdin = strings.NewReader(request.Script)

	// output = stdout + stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				return string(output), uint8(status.ExitStatus()), err
			}
		}
		return string(output), 0, err
	}

	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		return string(output), uint8(status.ExitStatus()), nil
	}

	return string(output), 0, nil
}
