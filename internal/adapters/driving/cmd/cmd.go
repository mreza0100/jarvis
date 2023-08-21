package cmd

import (
	"github.com/mreza0100/gptjarvis/internal/ports/cmdport"
	"github.com/mreza0100/gptjarvis/internal/ports/srvport"
	"github.com/urfave/cli/v2"
)

type cmd struct {
	srv *srvport.Services
}

type CmdHandlerParams struct {
	Srv *srvport.Services
}

func NewCMDHandler(params *CmdHandlerParams) cmdport.CMD {
	return &cmd{
		srv: params.Srv,
	}
}

func (c *cmd) Boot(ctx *cli.Context) error {
	err := c.srv.BootService.Start("jarvis.gpt")
	if err != nil {
		return err
	}

	return err
}
