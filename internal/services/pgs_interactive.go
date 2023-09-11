package services

import (
	"fmt"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runner_port"
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	openai "github.com/sashabaranov/go-openai"
)

type pgsInteractiveSrv struct {
	clinet             *openai.Client
	scriptCrashedTimes int

	runner     runnerport.PgsRunner
	interactor interactorport.Interactor
	history    historyport.History
	chat       chatport.Chat
}

func NewPgsSrv(req *srvport.PgsServicesReq) srvport.PgsInteractiveService {
	return &pgsInteractiveSrv{
		clinet:             openai.NewClient("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"),
		scriptCrashedTimes: 0,

		runner:     req.Runner,
		chat:       req.Chat,
		interactor: req.Interactor,
		history:    req.History,
	}
}

func (b *pgsInteractiveSrv) initLLMRole(prePrompt string) (*models.PgsReply, error) {
	prompt := &models.OSPrompt{
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

func (b *pgsInteractiveSrv) Start(prePrompt string) (err error) {
	reply, err := b.initLLMRole(prePrompt)
	if err != nil {
		return err
	}

	for {
		b.history.SaveResponse(reply)

		prompt := &models.PgsPrompt{}

		if reply.ReplyToUser != "" {
			b.interactor.Message(reply.ReplyToUser, reply.TokensUsed)
		}

		if reply.QueryRequest != nil {
			prompt.LastQueryResult, err = b.processScript(reply.QueryRequest)
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
			return err
		}

		b.interactor.Reply(reply)
	}
}

func (b *pgsInteractiveSrv) processScript(scriptRequest *models.PgsRunnerRequest) (*models.PgsScriptResult, error) {
	b.interactor.Script(scriptRequest)
	defer func() { b.scriptCrashedTimes = 0 }()

	runnerResult, err := b.runner.ExecScript(scriptRequest)
	if err != nil {
		b.scriptCrashedTimes++
		crashPrompt := "last executed command crashed. recovering... try again"
		reply := new(models.PgsReply)
		err = b.chat.Prompt(&models.PgsPrompt{
			ClientPrompt: &crashPrompt,
			UserPrompt:   nil,
			LastQueryResult: &models.PgsScriptResult{
				RunnerPgsResult: runnerResult,
			},
		}, reply)
		if err != nil {
			if b.scriptCrashedTimes > 5 {
				return nil, err
			}
			if reply != nil && reply.QueryRequest != nil {
				return b.processScript(scriptRequest)
			}
			return nil, err
		}
	}

	if runnerResult != nil {
		res := runnerResult.QueryResponses
		for _, r := range res {
			fmt.Printf("%+v\n", r)
		}
		fmt.Printf("runnerResult%+v\n", runnerResult.QueryResponses)
	}

	scriptResults := &models.PgsScriptResult{
		RunnerPgsResult: runnerResult,
	}
	b.interactor.ScriptResults(scriptResults)

	return scriptResults, nil
}
