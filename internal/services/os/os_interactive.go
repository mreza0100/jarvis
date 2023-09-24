package os_srvice

import (
	"fmt"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"

	openai "github.com/sashabaranov/go-openai"

	"github.com/mreza0100/jarvis/internal/ports/srvport"
)

type osService struct {
	clinet             *openai.Client
	scriptCrashedTimes int

	ConfigProvider cfgport.CfgProvider
	runner         runnerport.OSRunner
	chat           chatport.Chat
	interactor     interactorport.Interactor
	history        historyport.History
}

func NewOSService(req *srvport.OSServiceReq) srvport.OSService {
	return &osService{
		clinet:             openai.NewClient(req.ConfigProvider.GetConfigs().Token),
		scriptCrashedTimes: 0,

		ConfigProvider: req.ConfigProvider,
		runner:         req.Runner,
		chat:           req.Chat,
		interactor:     req.Interactor,
		history:        req.History,
	}
}

func (b *osService) initiateChat() (*models.OSReply, error) {
	prePrompt, err := b.ConfigProvider.LoadSavedFile("os.gpt")
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
	b.interactor.Reply(reply)

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

		prompt := &models.OSPrompt{}

		if reply.ReplyToUser != "" {
			b.interactor.Message(reply.ReplyToUser, b.chat.CountTokens())
		}

		if reply.ScriptRequest != nil {
			prompt.LastScriptResult, err = b.executeReplyScript(reply)
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

	SendPrompt:
		b.history.SavePrompt(prompt)
		if err := b.chat.Prompt(prompt, reply); err != nil {
			b.interactor.Error(err)
			clientErrReport := "Client error report: failed to process reply. Error:" + err.Error()
			prompt.ClientPrompt = &clientErrReport
			goto SendPrompt
		}
		b.interactor.Reply(reply)
	}
}

func (b *osService) executeReplyScript(reply *models.OSReply) (*models.OSScriptResult, error) {
	b.interactor.Script(reply.ScriptRequest)

	result, err := b.runner.ExecScript(reply.ScriptRequest)
	if err != nil {
		return nil, err
	}

	scriptResults := &models.OSScriptResult{RunnerOSResult: result}

	b.interactor.ScriptResults(scriptResults)

	return scriptResults, nil
}
