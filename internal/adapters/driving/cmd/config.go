package cmd

import (
	"github.com/mreza0100/jarvis/internal/adapters/driven/config"
	"github.com/mreza0100/jarvis/internal/services/config_service"
	"github.com/urfave/cli/v2"
)

func (c *cmd) BootstrapConfig(ctx *cli.Context) error {
	cfgProvider := config.NewConfigProvider()

	configDomain := config_service.NewConfigController(cfgProvider)

	trunk := ctx.Bool("trunk")
	return configDomain.Bootstrap(trunk)
}
