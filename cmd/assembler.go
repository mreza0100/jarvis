package cmd

import (
	"github.com/mreza0100/gptjarvis/internal/adapters/driven/chat"
	"github.com/mreza0100/gptjarvis/internal/adapters/driven/config"
	"github.com/mreza0100/gptjarvis/internal/adapters/driven/runner"
	"github.com/mreza0100/gptjarvis/internal/adapters/driving/cmd"
	"github.com/mreza0100/gptjarvis/internal/ports/cmdport"
	"github.com/mreza0100/gptjarvis/internal/ports/srvport"
	"github.com/mreza0100/gptjarvis/internal/services"
	"github.com/sashabaranov/go-openai"
)

func getCMDHandlers() cmdport.CMD {
	cfgProvider := config.NewConfigProvider()
	chat := chat.NewChat(&chat.NewChatReq{
		Clinet: openai.NewClient("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"),
	})

	services := services.NewServices(&srvport.ServicesReq{
		CfgProvider: cfgProvider,
		Chat:        chat,
		Runner:      runner.NewRunner(),
	})

	handlers := cmd.NewCMDHandler(&cmd.CmdHandlerParams{
		Srv: services,
	})

	return handlers
}
