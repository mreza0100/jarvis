package cmd

import (
	"encoding/json"
	"os"

	"github.com/mreza0100/jarvis/internal/adapters/driven/chat"
	"github.com/mreza0100/jarvis/internal/adapters/driven/config"
	"github.com/mreza0100/jarvis/internal/adapters/driven/history"
	"github.com/mreza0100/jarvis/internal/adapters/driven/interactor"
	"github.com/mreza0100/jarvis/internal/adapters/driven/runners"
	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	"github.com/mreza0100/jarvis/internal/services"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

func readPostgresConfig(path string) (*models.PostgresConnConfig, error) {
	rawContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	configs := new(models.PostgresConnConfig)
	return configs, json.Unmarshal(rawContent, configs)
}

func (c *cmd) PgsController(ctx *cli.Context) error {
	configFilePath := ctx.Args().Get(0)
	configs, err := readPostgresConfig(configFilePath)
	if err != nil {
		return err
	}

	cfgProvider := config.NewConfigProvider()
	history := history.NewHistory(cfgProvider)
	runner := runners.NewPgsRunner(&runners.PgsRunnerReq{
		Configs: configs,
	})

	chat := chat.NewChat(&chat.NewChatReq{
		Clinet: openai.NewClient("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"),
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
