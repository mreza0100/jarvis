package boot

import (
	"fmt"

	"github.com/mreza0100/gptjarvis/internal/models"
	"github.com/mreza0100/gptjarvis/internal/ports/chatport"
	"github.com/mreza0100/gptjarvis/internal/ports/historyport"
	"github.com/mreza0100/gptjarvis/internal/ports/interactorport"
	"github.com/mreza0100/gptjarvis/internal/ports/runnerport"
	modelstore "github.com/mreza0100/gptjarvis/models"

	openai "github.com/sashabaranov/go-openai"

	"github.com/mreza0100/gptjarvis/internal/ports/srvport"
)

type boot struct {
	clinet             *openai.Client
	scriptCrashedTimes int

	chat       chatport.Chat
	runner     runnerport.Runner
	interactor interactorport.Interactor
	history    historyport.History
}

func NewBootSrv(req *srvport.ServicesReq) srvport.BootService {
	return &boot{
		clinet:             openai.NewClient("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"),
		scriptCrashedTimes: 0,

		chat:       req.Chat,
		runner:     req.Runner,
		interactor: req.Interactor,
		history:    req.History,
	}
}

func (b *boot) initLLMRole(modelName string) (*models.Response, error) {
	content, err := modelstore.ModelsFS.ReadFile(modelName)
	if err != nil {
		return nil, err
	}
	modelDescriptor := string(content)

	prompt := &models.Prompt{
		ClientPrompt: &modelDescriptor,
	}
	b.history.SavePrompt(prompt)
	response, err := b.chat.Prompt(prompt)
	if err != nil {
		return nil, err
	}
	b.interactor.Response(response)

	return response, nil
}

func (b *boot) runScript(response *models.Response) (*models.ScriptResult, error) {
	b.interactor.Script(response.ScriptRequest)
	defer func() { b.scriptCrashedTimes = 0 }()

	output, status, err := b.runner.ExecuteScript(response.ScriptRequest)
	if err != nil {
		b.scriptCrashedTimes++
		crashPrompt := "last executed command crashed. recovering... try again"
		response, err = b.chat.Prompt(&models.Prompt{
			ClientPrompt: &crashPrompt,
			UserPrompt:   nil,
			LastExecutedScript: &models.ScriptResult{
				Stdout:         output,
				TerminalStatus: status,
			},
		})
		if err != nil {
			return b.runScript(response)
		}
	}

	scriptResults := &models.ScriptResult{
		Stdout:         output,
		TerminalStatus: status,
	}
	b.interactor.ScriptResults(scriptResults)

	return scriptResults, nil
}

func (b *boot) Start(modelName string) (err error) {
	defer func() {
		fmt.Println("Start Defer, err:", err)
	}()
	response, err := b.initLLMRole(modelName)
	if err != nil {
		return err
	}

	for {
		b.history.SaveResponse(response)
		prompt := &models.Prompt{}

		if response.MessageToUser != "" {
			b.interactor.Message(response.MessageToUser)
		}

		if response.ScriptRequest != nil {
			scriptResult, err := b.runScript(response)
			if err != nil {
				return err
			}
			if response.ScriptRequest.ReturnResults {
				prompt.LastExecutedScript = scriptResult
			}
		}
		if response.WaitForUserPrompt {
			userPrompt, err := b.interactor.GetUserInput()
			prompt.UserPrompt = &userPrompt
			if err != nil {
				return err
			}
		}

		b.history.SavePrompt(prompt)
		response, err = b.chat.Prompt(prompt)
		if err != nil {
			return err
		}
		b.interactor.Response(response)
	}
}
