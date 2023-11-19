package cmd

import (
	"github.com/mreza0100/jarvis/internal/adapters/driven/config"
	"github.com/mreza0100/jarvis/internal/services/configs"
	"github.com/urfave/cli/v2"
)

func (c *cmd) BootstrapConfig(ctx *cli.Context) error {
	cfgProvider := config.NewConfigProvider()

	configDomain := configs.NewConfigController(cfgProvider)

	trunk := ctx.Bool("trunk")
	return configDomain.Bootstrap(trunk)
}
