package cmd

import (
	"os"

	"github.com/mreza0100/jarvis/internal/adapters/driven/chat"
	"github.com/mreza0100/jarvis/internal/adapters/driven/config"
	"github.com/mreza0100/jarvis/internal/adapters/driven/history"
	"github.com/mreza0100/jarvis/internal/adapters/driven/runners"
	"github.com/mreza0100/jarvis/internal/adapters/driven/terminal"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	interactive_service "github.com/mreza0100/jarvis/internal/services/interactive"
	"github.com/urfave/cli/v2"
)

func (c *cmd) PgsController(ctx *cli.Context) error {
	cfgProvider := config.NewConfigProvider()
	configs := cfgProvider.LoadConfig()

	history := history.NewHistory(cfgProvider)

	runner := runners.NewPgsRunner(&runners.PgsRunnerReq{
		Configs: configs.ConfigFile.Postgres.PostgresConnConfig,
	})

	chat := chat.NewChat(&chat.NewChatReq{
		ChatConfigs: configs.ConfigFile.Postgres.Config,
	})
	terminal := terminal.NewTerminal(terminal.TerminalReq{
		CfgProvider: cfgProvider,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})

	postgresService := interactive_service.NewPgsService(&srvport.PgsServiceReq{
		ConfigProvider: cfgProvider,

		Chat:     chat,
		Runner:   runner,
		Terminal: terminal,
		History:  history,
	})

	return postgresService.RunInteractiveChat()
}
