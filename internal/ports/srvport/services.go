package srvport

import (
	"github.com/mreza0100/gptjarvis/internal/ports/cfgport"
	"github.com/mreza0100/gptjarvis/internal/ports/chatport"
	"github.com/mreza0100/gptjarvis/internal/ports/runnerport"
)

type ServicesReq struct {
	CfgProvider cfgport.CfgProvider
	Chat        chatport.Chat
	Runner      runnerport.Runner
}

type Services struct {
	BootService BootService
}

type BootService interface {
	Start() error
}
