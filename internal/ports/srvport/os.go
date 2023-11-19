package srvport

import (
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
	"github.com/mreza0100/jarvis/internal/ports/terminalport"
)

type OSServiceReq struct {
	ConfigProvider cfgport.CfgProvider
	Runner         runnerport.OSRunner
	Chat           chatport.Chat
	Terminal       terminalport.Terminal
	History        historyport.History
}

type OSServices struct {
	OSService OSService
}

type OSService interface {
	RunInteractiveChat() error
}
