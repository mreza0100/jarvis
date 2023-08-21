package boot

import (
	"fmt"
	"os"

	"github.com/mreza0100/gptjarvis/internal/models"
	"github.com/mreza0100/gptjarvis/internal/ports/chatport"
	"github.com/mreza0100/gptjarvis/internal/ports/interactorport"
	"github.com/mreza0100/gptjarvis/internal/ports/runnerport"

	openai "github.com/sashabaranov/go-openai"

	"github.com/mreza0100/gptjarvis/internal/ports/srvport"
)

type boot struct {
	clinet             *openai.Client
	chat               chatport.Chat
	runner             runnerport.Runner
	interactor         interactorport.Interactor
	scriptCrashedTimes int
}

func NewBootSrv(req *srvport.ServicesReq) srvport.BootService {
	return &boot{
		clinet:             openai.NewClient("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"),
		chat:               req.Chat,
		runner:             req.Runner,
		interactor:         req.Interactor,
		scriptCrashedTimes: 0,
	}
}

func (b *boot) initLLMRole() (*models.Response, error) {
	roleDescription, err := os.ReadFile("./configs/role-description.gpt")
	if err != nil {
		return nil, err
	}
	roleDescriptionStr := string(roleDescription)
	response, err := b.chat.Prompt(&models.Prompt{
		ClientPrompt: &roleDescriptionStr,
	})
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

func (b *boot) Start() (err error) {
	defer func() {
		fmt.Println("Start Defer, err:", err)
	}()
	response, err := b.initLLMRole()
	if err != nil {
		return err
	}

	for {
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

		response, err = b.chat.Prompt(prompt)
		if err != nil {
			return err
		}
		b.interactor.Response(response)
	}
}
