package services

import (
	"fmt"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runner_port"

	openai "github.com/sashabaranov/go-openai"

	"github.com/mreza0100/jarvis/internal/ports/srvport"
)

type osInteractiveSrv struct {
	clinet             *openai.Client
	scriptCrashedTimes int

	runner     runnerport.OSRunner
	chat       chatport.Chat
	interactor interactorport.Interactor
	history    historyport.History
}

func NewOSSrv(req *srvport.OSServicesReq) srvport.OSInteractiveService {
	return &osInteractiveSrv{
		clinet:             openai.NewClient("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"),
		scriptCrashedTimes: 0,

		runner:     req.Runner,
		chat:       req.Chat,
		interactor: req.Interactor,
		history:    req.History,
	}
}

func (b *osInteractiveSrv) initLLMRole(prePrompt string) (*models.OSReply, error) {
	prompt := &models.OSPrompt{
		ClientPrompt: &prePrompt,
	}
	b.history.SavePrompt(prompt)
	reply := new(models.OSReply)
	err := b.chat.Prompt(prompt, reply)
	if err != nil {
		return nil, err
	}
	b.interactor.Reply(reply)

	return reply, nil
}

func (b *osInteractiveSrv) Start(prePrompt string) (err error) {
	defer func() { fmt.Println("Start Defer, err:", err) }()

	reply, err := b.initLLMRole(prePrompt)
	if err != nil {
		return err
	}

	for {
		b.history.SaveReply(reply)

		prompt := &models.OSPrompt{}

		if reply.MessageToUser != "" {
			b.interactor.Message(reply.MessageToUser, reply.TokensUsed)
		}

		if reply.ScriptRequest != nil {
			prompt.LastScriptResult, err = b.processScript(reply)
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

func (b *osInteractiveSrv) processScript(reply *models.OSReply) (*models.OSScriptResult, error) {
	b.interactor.Script(reply.ScriptRequest)
	defer func() { b.scriptCrashedTimes = 0 }()

	result, err := b.runner.ExecScript(reply.ScriptRequest)
	if err != nil {
		b.scriptCrashedTimes++
		crashPrompt := "last executed command crashed. recovering... try again"
		reply := new(models.OSReply)
		err = b.chat.Prompt(&models.OSPrompt{
			ClientPrompt: &crashPrompt,
			UserPrompt:   nil,
			LastScriptResult: &models.OSScriptResult{
				RunnerOSResult: &models.OSRunnerResult{
					Stdout:     result.Stdout,
					StatusCode: result.StatusCode,
				},
			},
		}, reply)
		if err != nil {
			if b.scriptCrashedTimes > 5 {
				return nil, err
			}
			if reply != nil && reply.ScriptRequest != nil {
				return b.processScript(reply)
			}
			return nil, err
		}
	}

	scriptResults := &models.OSScriptResult{
		RunnerOSResult: &models.OSRunnerResult{
			Stdout:     result.Stdout,
			StatusCode: result.StatusCode,
		},
	}
	b.interactor.ScriptResults(scriptResults)

	return scriptResults, nil
}
