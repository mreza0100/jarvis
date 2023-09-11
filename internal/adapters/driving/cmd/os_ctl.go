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

func (c *cmd) OSController(_ *cli.Context) error {
	cfgProvider := config.NewConfigProvider()
	history := history.NewHistory(cfgProvider)
	runner := runners.NewOSRunner()

	chat := chat.NewChat(&chat.NewChatReq{
		Clinet: openai.NewClient("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"),
	})
	interactor := interactor.NewInteractor(interactor.InteractorArg{
		CfgProvider: cfgProvider,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	bootSrv := services.NewOSSrv(&srvport.OSServicesReq{
		Runner:     runner,
		Chat:       chat,
		Interactor: interactor,
		History:    history,
	})

	template, err := c.getTemplate("os.gpt")
	if err != nil {
		return err
	}

	return bootSrv.Start(template)
}
