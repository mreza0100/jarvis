package services

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/mreza0100/jarvis/internal/ports/chatport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"

	openai "github.com/sashabaranov/go-openai"

	"github.com/mreza0100/jarvis/internal/ports/srvport"
)

type OSReply struct {
	MessageToUser     string           `json:"messageToUser"`
	WaitForUserPrompt bool             `json:"waitForUserPrompt"`
	ScriptRequest     *RunnerOSRequest `json:"Script"`
	TokensUsed        int              `json:"_"`
}

type RunnerOSRequest struct {
	Runtime string `json:"Runtime"`
	Script  string `json:"Script"`
}

type RunnerOSResponse struct {
	Stdout     string `json:"Stdout"`
	StatusCode int    `json:"StatusCode"`
}

type Prompt struct {
	ClientPrompt     *string       `json:"ClientPrompt"`
	UserPrompt       *string       `json:"UserPrompt"`
	Screen           *Screen       `json:"Screen"`
	LastScriptResult *ScriptResult `json:"LastScriptResult"`
}

type Screen struct {
	Width  int `json:"Width"`
	Height int `json:"Height"`
}

type ScriptResult struct {
	RunnerOSResult *RunnerOSResponse `json:"RunnerOSResult"`
}

// ---

type osInteractiveSrv struct {
	clinet             *openai.Client
	scriptCrashedTimes int

	chat       chatport.Chat
	interactor interactorport.Interactor
	history    historyport.History
}

func NewOSSrv(req *srvport.ServicesReq) srvport.BootService {
	return &osInteractiveSrv{
		clinet:             openai.NewClient("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"),
		scriptCrashedTimes: 0,

		chat:       req.Chat,
		interactor: req.Interactor,
		history:    req.History,
	}
}

func (b *osInteractiveSrv) initLLMRole(prePrompt string) (*OSReply, error) {
	prompt := &Prompt{
		ClientPrompt: &prePrompt,
	}
	b.history.SavePrompt(prompt)
	reply := new(OSReply)
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
		b.history.SaveResponse(reply)

		prompt := &Prompt{}

		if reply.MessageToUser != "" {
			b.interactor.Message(reply.MessageToUser, reply.TokensUsed)
		}

		if reply.ScriptRequest != nil {
			scriptResult, err := b.processScript(reply)
			if err != nil {
				return err
			}
			prompt.LastScriptResult = scriptResult
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

func (b *osInteractiveSrv) processScript(reply *OSReply) (*ScriptResult, error) {
	b.interactor.Script(reply.ScriptRequest)
	defer func() { b.scriptCrashedTimes = 0 }()

	result, err := b.execScript(reply.ScriptRequest)
	if err != nil {
		b.scriptCrashedTimes++
		crashPrompt := "last executed command crashed. recovering... try again"
		reply := new(OSReply)
		err = b.chat.Prompt(&Prompt{
			ClientPrompt: &crashPrompt,
			UserPrompt:   nil,
			LastScriptResult: &ScriptResult{
				RunnerOSResult: &RunnerOSResponse{
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

	scriptResults := &ScriptResult{
		RunnerOSResult: &RunnerOSResponse{
			Stdout:     result.Stdout,
			StatusCode: result.StatusCode,
		},
	}
	b.interactor.ScriptResults(scriptResults)

	return scriptResults, nil
}

func (b *osInteractiveSrv) execScript(req *RunnerOSRequest) (*RunnerOSResponse, error) {
	cmd := exec.Command(req.Runtime)
	cmd.Stdin = strings.NewReader(req.Script)

	// output = stdout + stderr
	rawOutput, err := cmd.CombinedOutput()
	output := string(rawOutput)
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				return &RunnerOSResponse{
					StatusCode: status.ExitStatus(),
					Stdout:     output,
				}, err
			}
		}
		return &RunnerOSResponse{
			StatusCode: 0,
			Stdout:     output,
		}, err
	}

	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		return &RunnerOSResponse{
			StatusCode: status.ExitStatus(),
			Stdout:     output,
		}, nil
	}

	return &RunnerOSResponse{
		StatusCode: 0,
		Stdout:     output,
	}, nil
}
