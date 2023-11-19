package cmd

import (
	"os"

	"github.com/mreza0100/jarvis/internal/adapters/driven/chat"
	"github.com/mreza0100/jarvis/internal/adapters/driven/config"
	"github.com/mreza0100/jarvis/internal/adapters/driven/history"
	"github.com/mreza0100/jarvis/internal/adapters/driven/interactor"
	"github.com/mreza0100/jarvis/internal/adapters/driven/runners"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	os_srvice "github.com/mreza0100/jarvis/internal/services/os"
	"github.com/urfave/cli/v2"
)

func (c *cmd) OSController(ctx *cli.Context) error {
	cfgProvider := config.NewConfigProvider()
	configs := cfgProvider.LoadConfig()

	history := history.NewHistory(cfgProvider)

	runner := runners.NewOSRunner()
	chat := chat.NewChat(&chat.NewChatReq{
		ChatConfigs: configs.ConfigFile.OS.Config,
	})
	interactor := interactor.NewInteractor(interactor.InteractorReq{
		CfgProvider: cfgProvider,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	bootSrv := os_srvice.NewOSService(&srvport.OSServiceReq{
		ConfigProvider: cfgProvider,
		Runner:         runner,
		Chat:           chat,
		Interactor:     interactor,
		History:        history,
	})

	return bootSrv.RunInteractiveChat()
}
