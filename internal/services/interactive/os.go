package interactive_service

import (
	"fmt"
	"time"

	"github.com/mreza0100/jarvis/internal/models"
	errs "github.com/mreza0100/jarvis/internal/pkg/errors"
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
	terminal       terminalport.Terminal
	history        historyport.History
}

func NewOSService(req *srvport.OSServiceReq) srvport.OSService {
	return &osService{
		scriptCrashedTimes: 0,

		Screen:         &models.Screen{},
		ConfigProvider: req.ConfigProvider,
		runner:         req.Runner,
		chat:           req.Chat,
		terminal:       req.Terminal,
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
			e := &errs.Error{}
			if errors.As(err, e) {
				switch e.Code {
				case errs.API_AUTH:
					return err
				case errs.API_INTERNAL_ERROR:
					b.terminal.Error(err)
					time.Sleep(time.Second * 3)
					goto SendPrompt
				case errs.API_RATE_LIMIT:
					waitDuration, ok := e.Data["wait"].(time.Duration)
					if !ok {
						return errors.New("cast failed, e.Params[1].(time.Duration)")
					}
					b.terminal.Error(errors.Wrapf(err, "retrying in %s", waitDuration.String()))
					time.Sleep(waitDuration)
					goto SendPrompt
				}
			}

			clientErrReport := errors.Wrap(err, "Client error report: failed to process reply")
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
