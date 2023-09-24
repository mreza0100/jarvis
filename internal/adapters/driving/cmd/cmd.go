package cmd

import (
	"github.com/mreza0100/jarvis/internal/ports/cmdport"
)

type cmd struct{}

type CmdHandlerParams struct{}

func NewCMDHandler(params *CmdHandlerParams) cmdport.CMD {
	return &cmd{}
}
