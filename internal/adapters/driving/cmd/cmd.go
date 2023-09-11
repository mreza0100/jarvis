package cmd

import (
	"github.com/mreza0100/jarvis/internal/ports/cmdport"
	templatestore "github.com/mreza0100/jarvis/templates"
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
