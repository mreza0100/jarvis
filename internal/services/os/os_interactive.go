package os_srvice

import (
	"fmt"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
	"github.com/mreza0100/jarvis/internal/ports/terminalport"
	"github.com/pkg/errors"

	"github.com/mreza0100/jarvis/internal/ports/srvport"
)

type osService struct {
	scriptCrashedTimes int

	Screen         *models.Screen
	ConfigProvider cfgport.CfgProvider
	runner         runnerport.OSRunner
	chat           chatport.Chat
	terminal       terminalport.Interactor
	history        historyport.History
}

func NewOSService(req *srvport.OSServiceReq) srvport.OSService {
	return &osService{
		scriptCrashedTimes: 0,

		Screen:         &models.Screen{},
		ConfigProvider: req.ConfigProvider,
		runner:         req.Runner,
		chat:           req.Chat,
		terminal:       req.Interactor,
		history:        req.History,
	}
}

func (b *osService) initiateChat() (*models.OSReply, error) {
	prePrompt, err := b.ConfigProvider.LoadStoredFile("os.gpt")
	if err != nil {
		return nil, err
	}
	prompt := &models.OSPrompt{
		ClientPrompt: &prePrompt,
	}
	b.history.SavePrompt(prompt)

	reply := new(models.OSReply)
	if err := b.chat.Prompt(prompt, reply); err != nil {
		return nil, err
	}
	b.terminal.Reply(reply)

	return reply, nil
}

func (b *osService) RunInteractiveChat() (err error) {
	defer func() { fmt.Println("RunInteractiveChat Defer, err:", err) }()

	reply, err := b.initiateChat()
	if err != nil {
		return err
	}

	for {
		b.history.SaveReply(reply)

		prompt := &models.OSPrompt{Screen: b.Screen.GetScreen()}

		if reply.ReplyToUser != "" {
			b.terminal.PrintReply(reply.ReplyToUser, b.chat.GetRateLimitInsights())
		}

		if reply.ScriptRequest != nil {
			prompt.LastScriptResult, err = b.executeReplyScript(reply)
			if err != nil {
				clientPrompt := "Error detected"
				prompt = &models.OSPrompt{
					ClientPrompt: &clientPrompt,
					LastScriptResult: &models.OSScriptResult{
						RunnerOSResult: &models.OSRunnerResult{Stderr: err.Error(), StatusCode: 1},
					},
				}
				goto SendPrompt
			}
		}
		if reply.WaitForUserPrompt {
			userPrompt, err := b.terminal.GetUserInput()
			prompt.UserPrompt = &userPrompt
			if err != nil {
				return err
			}
		}

	SendPrompt:
		b.history.SavePrompt(prompt)
		if err := b.chat.Prompt(prompt, reply); err != nil {
			clientErrReport := errors.Wrap(err, "Client error report: failed to process reply. Error:")
			clientErrReportStr := clientErrReport.Error()
			b.terminal.Error(clientErrReport)
			prompt.ClientPrompt = &clientErrReportStr
			goto SendPrompt
		}
		b.terminal.Reply(reply)
	}
}

func (b *osService) executeReplyScript(reply *models.OSReply) (*models.OSScriptResult, error) {
	b.terminal.Script(reply.ScriptRequest)

	result, err := b.runner.ExecScript(reply.ScriptRequest)
	if err != nil {
		return nil, err
	}

	scriptResults := &models.OSScriptResult{RunnerOSResult: result}

	b.terminal.ScriptResults(scriptResults)

	return scriptResults, nil
}
