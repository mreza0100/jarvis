package cmd

import (
	"os"

	"github.com/mreza0100/jarvis/internal/adapters/driven/chat"
	"github.com/mreza0100/jarvis/internal/adapters/driven/config"
	"github.com/mreza0100/jarvis/internal/adapters/driven/history"
	"github.com/mreza0100/jarvis/internal/adapters/driven/interactor"
	"github.com/mreza0100/jarvis/internal/adapters/driven/runners"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	"github.com/mreza0100/jarvis/internal/services"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

func (c *cmd) PgsController(ctx *cli.Context) error {
	configFilePath := ctx.Args().Get(0)

	cfgProvider := config.NewConfigProvider(&configFilePath)
	history := history.NewHistory(cfgProvider)
	cfg := cfgProvider.GetCfg()
	runner := runners.NewPgsRunner(&runners.PgsRunnerReq{
		Configs: &cfg.ConfigFile.PostgresConfig.PostgresConnConfig,
	})

	chat := chat.NewChat(&chat.NewChatReq{
		Clinet: openai.NewClient(cfg.Token),
	})
	interactor := interactor.NewInteractor(interactor.InteractorArg{
		CfgProvider: cfgProvider,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	bootSrv := services.NewPgsSrv(&srvport.PgsServicesReq{
		Chat:       chat,
		Runner:     runner,
		Interactor: interactor,
		History:    history,
	})

	template, err := c.getTemplate("postgres.gpt")
	if err != nil {
		return err
	}

	return bootSrv.Start(template)
}
