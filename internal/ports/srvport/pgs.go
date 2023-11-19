package srvport

import (
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
	"github.com/mreza0100/jarvis/internal/ports/terminalport"
)

type PgsServiceReq struct {
	ConfigProvider cfgport.CfgProvider
	Runner         runnerport.PgsRunner
	Chat           chatport.Chat
	Interactor     terminalport.Interactor
	History        historyport.History
}

type PgsServices struct {
	PgsService PgsService
}

type PgsService interface {
	RunInteractiveChat() error
}
