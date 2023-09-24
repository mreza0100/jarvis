package cmd

import (
	"os"

	"github.com/mreza0100/jarvis/internal/adapters/driven/chat"
	"github.com/mreza0100/jarvis/internal/adapters/driven/config"
	"github.com/mreza0100/jarvis/internal/adapters/driven/history"
	"github.com/mreza0100/jarvis/internal/adapters/driven/interactor"
	"github.com/mreza0100/jarvis/internal/adapters/driven/runners"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	pgs_srvice "github.com/mreza0100/jarvis/internal/services/pgs"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

func (c *cmd) PgsController(ctx *cli.Context) error {
	configFilePath := ctx.Args().Get(0)
	cfgProvider := config.NewConfigProvider(&configFilePath)
	configs := cfgProvider.GetConfigs()

	history := history.NewHistory(cfgProvider)

	runner := runners.NewPgsRunner(&runners.PgsRunnerReq{
		Configs: &configs.ConfigFile.PostgresConfig.PostgresConnConfig,
	})

	chat := chat.NewChat(&chat.NewChatReq{
		Clinet: openai.NewClient(configs.Token),
	})
	interactor := interactor.NewInteractor(interactor.InteractorReq{
		CfgProvider: cfgProvider,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})

	postgresService := pgs_srvice.NewPgsService(&srvport.PgsServiceReq{
		ConfigProvider: cfgProvider,

		Chat:       chat,
		Runner:     runner,
		Interactor: interactor,
		History:    history,
	})

	return postgresService.RunInteractiveChat()
}
