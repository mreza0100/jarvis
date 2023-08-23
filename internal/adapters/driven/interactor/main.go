package interactor

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"

	"github.com/mreza0100/gptjarvis/internal/models"
	"github.com/mreza0100/gptjarvis/internal/pkg/terminal"
	"github.com/mreza0100/gptjarvis/internal/ports/cfgport"
	"github.com/mreza0100/gptjarvis/internal/ports/interactorport"
)

type interactor struct {
	mode models.Mode

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

type InteractorArg struct {
	CfgProvider cfgport.CfgProvider

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewInteractor(args InteractorArg) interactorport.Interactor {
	return &interactor{
		mode: args.CfgProvider.GetCfg().Mode,

		stdin:  args.Stdin,
		stdout: args.Stdout,
		stderr: args.Stderr,
	}
}

func (i *interactor) GetUserInput() (string, error) {
	fmt.Print("\n==> ")
	reader := bufio.NewReader(i.stdin)
	userPrompt, _ := reader.ReadString('\n')
	return strings.TrimSpace(userPrompt), nil
}

func (i *interactor) setColor(color color) {
	fmt.Print(color)
}

func (i *interactor) unsetColor() {
	fmt.Printf("%s[%dm", "\x1b", 0)
}

func (i *interactor) log(color color, title string, message any) {
	i.setColor(color)
	defer i.unsetColor()

	fmt.Printf("%s%+v\n", title, message)
}

func (i *interactor) Message(message string, usedTokens int) {
	screen, err := terminal.GetTerminalSize()
	if err != nil {
		log.Fatal(err)
		return
	}

	renderedMarkdown := markdown.Render(message, screen.Width, 0)
	i.log(colorYellow, "Used Tokens: ", usedTokens)
	i.log(colorCyan, "\n", string(renderedMarkdown))
}

func (i *interactor) Script(script *models.ScriptRequest) {
	i.setColor(colorRed)
	defer i.unsetColor()
	if i.mode.IsDev() {
		fmt.Printf("\n- Script:\n")
		fmt.Printf("Runtime=%s\n", script.Runtime)
		fmt.Printf("Script=%s\n", script.Script)
		fmt.Printf("ReturnResults=%+v\n", script.ReturnResults)
		fmt.Printf("\n--\n")
	}
}

func (i *interactor) ScriptResults(result *models.ScriptResult) {
	if i.mode.IsDev() {
		i.log(colorBlue, "ScriptResults", result)
	}
}

func (i *interactor) Response(response *models.Response) {
	i.setColor(colorPurple)
	defer i.unsetColor()
	if i.mode.IsDev() {
		fmt.Printf("\n--\n%s::\n", "Response")
		fmt.Printf("MessageToUser=%s\n", response.MessageToUser)
		fmt.Printf("WaitForUserPrompt=%v\n", response.WaitForUserPrompt)
		fmt.Printf("ScriptRequest=%+v\n", response.ScriptRequest)
		fmt.Printf("\n--\n")
	}
}
