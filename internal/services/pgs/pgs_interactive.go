package pgs_srvice

import (
	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	openai "github.com/sashabaranov/go-openai"
)

type pgsService struct {
	clinet *openai.Client

	ConfigProvider cfgport.CfgProvider
	runner         runnerport.PgsRunner
	interactor     interactorport.Interactor
	history        historyport.History
	chat           chatport.Chat
}

func NewPgsService(req *srvport.PgsServiceReq) srvport.PgsService {
	return &pgsService{
		clinet: openai.NewClient(req.ConfigProvider.GetConfigs().Token),

		ConfigProvider: req.ConfigProvider,
		runner:         req.Runner,
		chat:           req.Chat,
		interactor:     req.Interactor,
		history:        req.History,
	}
}

func (b *pgsService) initiateChat() (*models.PgsReply, error) {
	prePrompt, err := b.ConfigProvider.LoadSavedFile("postgres.gpt")
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

		prompt := &models.PgsPrompt{}

		if reply.ReplyToUser != "" {
			b.interactor.Message(reply.ReplyToUser, b.chat.CountTokens())
		}

		if reply.QueryRequest != nil {
			prompt.LastQueryResult, err = b.executeReplyQuery(reply.QueryRequest)
			if err != nil {
				return err
			}
		}
		if reply.WaitForUserPrompt {
			userPrompt, err := b.interactor.GetUserInput()
			prompt.UserPrompt = &userPrompt
			if err != nil {
				return err
			}
		}

		b.history.SavePrompt(prompt)
		if err := b.chat.Prompt(prompt, reply); err != nil {
			clientPrompt := "client crash, error: " + err.Error()
			prompt := &models.PgsPrompt{
				UserPrompt:   nil,
				ClientPrompt: &clientPrompt,
			}

			if err := b.chat.Prompt(prompt, reply); err != nil {
				return err
			}
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
