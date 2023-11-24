package interactive_service

import (
	"time"

	"github.com/mreza0100/jarvis/internal/models"
	errs "github.com/mreza0100/jarvis/internal/pkg/errors"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	"github.com/mreza0100/jarvis/internal/ports/terminalport"
	"github.com/pkg/errors"
)

type pgsService struct {
	Screen         *models.Screen
	ConfigProvider cfgport.CfgProvider
	runner         runnerport.PgsRunner
	terminal       terminalport.Terminal
	history        historyport.History
	chat           chatport.Chat
}

func NewPgsService(req *srvport.PgsServiceReq) srvport.PgsService {
	return &pgsService{
		Screen:         &models.Screen{},
		ConfigProvider: req.ConfigProvider,
		runner:         req.Runner,
		chat:           req.Chat,
		terminal:       req.Terminal,
		history:        req.History,
	}
}

// TODO: Refactor this and os, both RunInteractiveChat methods must connect to the same struct
// TODO: Also decuple blocks, break them into functions and set there data in the state of the struct, so the helper common methods in that struct can change the state of the loop
// TODO: State values: waitForUser, Executing and etc, also state needs to containt the prompt, answer and other data
func (b *pgsService) RunInteractiveChat() error {
	prePrompt, err := b.ConfigProvider.LoadStoredFile("postgres.gpt")
	if err != nil {
		return err
	}

	reply := new(models.PgsReply)
	prompt := &models.PgsPrompt{
		ClientPrompt: &prePrompt,
		Screen:       b.Screen.GetScreen(),
	}

	for {
		b.history.SavePrompt(prompt)

	SendPrompt:
		b.history.SavePrompt(prompt)
		if err := b.chat.Prompt(prompt, reply); err != nil {
			// CONTINUE WHERE WE LEFT OFF:
			// bring os and pgs domains next to each other so they can share some methods
			// then we need to take the repeated handleing codes to 1 method so they can handle chat errors in the same way
			// then we need to create garbage collector next to chat and automatically call it from here
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

			clientErrReport := errors.Wrap(err, "Client error report: failed to process reply or parse json, reason:")
			clientErrReportStr := clientErrReport.Error()
			b.terminal.Error(clientErrReport)
			prompt.ClientPrompt = &clientErrReportStr
			goto SendPrompt
		}
		b.history.SaveReply(reply)
		b.terminal.Reply(reply)

		if reply.ReplyToUser != "" {
			b.terminal.PrintReply(reply.ReplyToUser, b.chat.GetRateLimitInsights())
		}

		if reply.QueryRequest != nil {
			prompt.LastQueryResult, err = b.executeReplyQuery(reply.QueryRequest)
			if err != nil {
				clientPrompt := "Error detected"
				prompt = &models.PgsPrompt{
					ClientPrompt: &clientPrompt,
					LastQueryResult: &models.PgsScriptResult{
						RunnerPgsResult: &models.PgsRunnerResponse{Err: err},
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
	}
}

func (b *pgsService) executeReplyQuery(queryRequest *models.PgsRunnerRequest) (*models.PgsScriptResult, error) {
	b.terminal.Script(queryRequest)

	runnerResult, err := b.runner.ExecScript(queryRequest)
	if err != nil {
		return nil, err
	}

	scriptResults := &models.PgsScriptResult{
		RunnerPgsResult: runnerResult,
	}
	b.terminal.ScriptResults(scriptResults)

	return scriptResults, nil
}
