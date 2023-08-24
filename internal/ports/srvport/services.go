package srvport

import (
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	"github.com/mreza0100/jarvis/internal/ports/runnerport"
)

type ServicesReq struct {
	Chat       chatport.Chat
	Runner     runnerport.Runner
	Interactor interactorport.Interactor
	History    historyport.History
}

type Services struct {
	BootService BootService
}

type BootService interface {
	Start(modelName string) error
}
