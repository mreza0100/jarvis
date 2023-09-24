package srvport

import (
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
)

type OSServiceReq struct {
	ConfigProvider cfgport.CfgProvider
	Runner         runnerport.OSRunner
	Chat           chatport.Chat
	Interactor     interactorport.Interactor
	History        historyport.History
}

type OSServices struct {
	OSService OSService
}

type OSService interface {
	RunInteractiveChat() error
}
