package srvport

import (
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runner_port"
)

type PgsServicesReq struct {
	ConfigProvider cfgport.CfgProvider
	Runner         runnerport.PgsRunner
	Chat           chatport.Chat
	Interactor     interactorport.Interactor
	History        historyport.History
}

type PgsServices struct {
	BootService PgsInteractiveService
}

type PgsInteractiveService interface {
	Start(modelName string) error
}
