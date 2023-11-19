package pgs_srvice

import (
	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	"github.com/pkg/errors"
)

type pgsService struct {
	Screen         *models.Screen
	ConfigProvider cfgport.CfgProvider
	runner         runnerport.PgsRunner
	interactor     interactorport.Interactor
	history        historyport.History
	chat           chatport.Chat
}

func NewPgsService(req *srvport.PgsServiceReq) srvport.PgsService {
	return &pgsService{
		Screen:         &models.Screen{},
		ConfigProvider: req.ConfigProvider,
		runner:         req.Runner,
		chat:           req.Chat,
		interactor:     req.Interactor,
		history:        req.History,
	}
}

func (b *pgsService) initiateChat() (*models.PgsReply, error) {
	prePrompt, err := b.ConfigProvider.LoadStoredFile("postgres.gpt")
	if err != nil {
		return nil, err
	}
	prompt := &models.PgsPrompt{
		ClientPrompt: &prePrompt,
	}
	b.history.SavePrompt(prompt)
	reply := new(models.PgsReply)
	if err := b.chat.Prompt(prompt, reply); err != nil {
		return nil, err
	}
	b.interactor.Reply(reply)

	return reply, nil
}

func (b *pgsService) RunInteractiveChat() error {
	reply, err := b.initiateChat()
	if err != nil {
		return err
	}

	for {
		b.history.SaveReply(reply)

		prompt := &models.PgsPrompt{Screen: b.Screen.GetScreen()}

		if reply.ReplyToUser != "" {
			b.interactor.Message(reply.ReplyToUser, b.chat.CountTokens())
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
			userPrompt, err := b.interactor.GetUserInput()
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
			b.interactor.Error(clientErrReport)
			prompt.ClientPrompt = &clientErrReportStr
			goto SendPrompt
		}
		b.interactor.Reply(reply)
	}
}

func (b *pgsService) executeReplyQuery(queryRequest *models.PgsRunnerRequest) (*models.PgsScriptResult, error) {
	b.interactor.Script(queryRequest)

	runnerResult, err := b.runner.ExecScript(queryRequest)
	if err != nil {
		return nil, err
	}

	scriptResults := &models.PgsScriptResult{
		RunnerPgsResult: runnerResult,
	}
	b.interactor.ScriptResults(scriptResults)

	return scriptResults, nil
}
