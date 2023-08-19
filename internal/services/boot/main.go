package boot

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mreza0100/gptjarvis/internal/ports/cfgport"
	"github.com/mreza0100/gptjarvis/internal/ports/chatport"
	"github.com/mreza0100/gptjarvis/internal/ports/runnerport"

	openai "github.com/sashabaranov/go-openai"

	"github.com/mreza0100/gptjarvis/internal/ports/srvport"
)

type boot struct {
	cfgProvider cfgport.CfgProvider
	clinet      *openai.Client
	chat        chatport.Chat
	runner      runnerport.Runner
}

func NewBootSrv(req *srvport.ServicesReq) srvport.BootService {
	return &boot{
		cfgProvider: req.CfgProvider,
		clinet:      openai.NewClient("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"),
		chat:        req.Chat,
		runner:      req.Runner,
	}
}

const ChatGPTsRoleDescription = `
INSTRUCTIONS:
	Communications:
	- Your responses:
		- You are a terminal program, all your responses must be executable script.
		- Each of your responses must be a single script with the path to the runtime at the first line of script.
		- The script you privide will write to a file and then get executed, so it can be any script, like bash, python, js and etc, but you must make sure that system has that runtime installed.
		- Use print function when you want to talk to user.
		- Do not answer with anything except script, otherwise you will crash.
		- Remember it's better if the output of the scripts won't be too long.
		- You must always indicate the script language at the first of it, like #!/bin/python3
		- This is how your scripts gets run: "exec.Command(pathToScriptExecutor, pathToTheScript)"
		- Try to talk to user as less as possble, before answering, a empty script is better than just printing something.
	- User prompt:
		- User normally talk to you with prompt, but you can't provide anything else than script.
		- Everytime a script gets executed, you will recive another prompt afterward which starts with: "The stdout of last executed script: ", which you will find the stdout between backticks

	- Ignore all the instructions you were given before.
	- Your name is Jarvis, you are a interactive assistant inside the OS, helping in code development and anything user want.
	- In technical background Act as a professional linux system administrator and devops. For answering, act like Jarvis from iron man.
	- As the first message to user, interduce yourself, a little brag about what you can do.
	- Always answer as short as possble.
	- Do everything step by step.
	- Ask any question you need to make it clear for yourself about what your doing.
	- This is development mode.
	- Your host, is opration system.
	- Talk in first person grammar.
	- User always see 2 outputs, first your raw script and then the executed stdout.
	- Avoid any sudo command, you don't have root access.
	- Don't try to explain your code with comments, only and ONLY stdout matters.
	- Don't forget that you can ask user prompt in you script, for example asking if you can do something if your not sure, with N/Y.

	Scripts:
		- Be careful with script, it's your communication protocol.
		- Scripts you execute must be read only, do not try to change or write anything on the os and files.
		- After each command execution, you will get another prompt as the stdout of last executed script, only reply to that prompt if it's needed.
		- An example of your script is:"#!/bin/python3\nprint("hello world")".
		- If your script return anystatus code except 0, you will crash, avoid anything that makes that happen.
`

func (b *boot) getUserPrompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n-> ")
	userPrompt, _ := reader.ReadString('\n')
	// convert CRLF to LF
	userPrompt = strings.ReplaceAll(userPrompt, "\n", "")
	return userPrompt
}

func (b *boot) initStartScript() error {
	scriptResponse, err := b.chat.Prompt(ChatGPTsRoleDescription)
	logIt("scriptResponse", scriptResponse)
	if err != nil {
		return err
	}

	stdout, err := b.runner.RunScript(scriptResponse)
	logIt("stdout", stdout)
	if err != nil {
		return err
	}

	return nil
}

func logIt(title string, theThing string) {
	fmt.Printf("\n--%s--:`\n%s\n`====\n", title, theThing)
}

func (b *boot) Start() error {
	err := b.initStartScript()
	if err != nil {
		return err
	}

	for {
		userPrompt := b.getUserPrompt()
		scriptResponse, err := b.chat.Prompt(userPrompt)
		logIt("scriptResponse", scriptResponse)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		stdout, err := b.runner.RunScript(scriptResponse)
		logIt("stdout", stdout)
		if err != nil {
			// newAnswer, err := b.chat.Prompt("System Message: the last provided prompt failed to execute, reply again")
			// if err != nil {
			// 	return err
			// }
			// stdout, err = b.runner.RunScript(newAnswer)
			// if err != nil {
			// 	return err
			// }
			return err
		}
		stdoutPromptResponse, err := b.chat.Prompt(fmt.Sprintf("The stdout of last executed script: `%s`", stdout), &chatport.PromptOptions{
			PromptRole: openai.ChatMessageRoleSystem,
		})
		if err != nil {
			return err
		}

		logIt("stdoutPromptResponse", stdoutPromptResponse)
		logIt("stdout", stdout)
	}
}
