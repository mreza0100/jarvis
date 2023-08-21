package cmd

import (
	"os"

	"github.com/mreza0100/gptjarvis/internal/adapters/driven/chat"
	"github.com/mreza0100/gptjarvis/internal/adapters/driven/config"
	"github.com/mreza0100/gptjarvis/internal/adapters/driven/history"
	"github.com/mreza0100/gptjarvis/internal/adapters/driven/interactor"
	"github.com/mreza0100/gptjarvis/internal/adapters/driven/runner"
	"github.com/mreza0100/gptjarvis/internal/adapters/driving/cmd"
	"github.com/mreza0100/gptjarvis/internal/ports/cmdport"
	"github.com/mreza0100/gptjarvis/internal/ports/srvport"
	"github.com/mreza0100/gptjarvis/internal/services"
	"github.com/sashabaranov/go-openai"
)

func getCMDHandlers() cmdport.CMD {
	cfgProvider := config.NewConfigProvider()
	history := history.NewHistory(cfgProvider)
	runner := runner.NewRunner()

	chat := chat.NewChat(&chat.NewChatReq{
		Clinet: openai.NewClient("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"),
	})
	interactor := interactor.NewInteractor(interactor.InteractorArg{
		CfgProvider: cfgProvider,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})

	services := services.NewServices(&srvport.ServicesReq{
		Chat:       chat,
		Runner:     runner,
		Interactor: interactor,
		History:    history,
	})

	handlers := cmd.NewCMDHandler(&cmd.CmdHandlerParams{
		Srv: services,
	})

	return handlers
}
