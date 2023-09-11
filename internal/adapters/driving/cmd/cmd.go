package cmd

import (
	"errors"

	"github.com/mreza0100/jarvis/internal/ports/cmdport"
	templatestore "github.com/mreza0100/jarvis/templates"
	"github.com/urfave/cli/v2"
)

type cmd struct{}

type CmdHandlerParams struct{}

func NewCMDHandler(params *CmdHandlerParams) cmdport.CMD {
	return &cmd{}
}

func (c *cmd) getTemplate(templateName string) (string, error) {
	content, err := templatestore.ModelsFS.ReadFile(templateName)
	if err != nil {
		return "", err
	}

	templ := string(content)
	return templ, nil
}

func (c *cmd) Interactive(ctx *cli.Context) error {
	controllerName := ctx.Args().Get(0)

	switch controllerName {
	case "os":
		return c.OSController(ctx)
	case "postgres":
		return c.pgsController(ctx)

	default:
		return errors.New("Failed to find jarvis model")
	}
}
