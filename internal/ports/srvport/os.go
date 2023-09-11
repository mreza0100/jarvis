package srvport

import (
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runner_port"
)

type OSServicesReq struct {
	Runner     runnerport.OSRunner
	Chat       chatport.Chat
	Interactor interactorport.Interactor
	History    historyport.History
}

type OSServices struct {
	BootService OSInteractiveService
}

type OSInteractiveService interface {
	Start(modelName string) error
}