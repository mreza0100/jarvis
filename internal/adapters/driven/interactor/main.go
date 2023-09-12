package interactor

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/pkg/tools"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/interactorport"
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

	fmt.Printf("%s:\n%+v\n", title, message)
}

func (i *interactor) logToken(color color, tokens int) {
	i.setColor(color)
	defer i.unsetColor()

	fmt.Printf("Tokens Used: %d\n", tokens)
}

func (i *interactor) Message(message string, usedTokens int) {
	width, _, err := tools.GetTerminalSize()
	if err != nil {
		log.Fatal(err)
		return
	}

	renderedMarkdown := markdown.Render(message, width, 0)
	i.logToken(colorYellow, usedTokens)
	i.log(colorCyan, "\n", string(renderedMarkdown))
}

func (i *interactor) Script(script interface{}) {
	i.setColor(colorRed)
	defer i.unsetColor()
	if i.mode.IsDev() {
		fmt.Printf("\n- Script:\n")
		fmt.Printf("script.String()=%+v\n", script)
		fmt.Printf("\n--\n")
	}
}

func (i *interactor) ScriptResults(result interface{}) {
	if i.mode.IsDev() {
		i.log(colorBlue, "Script Results", result)
	}
}

func (i *interactor) Reply(reply interface{}) {
	i.setColor(colorPurple)
	defer i.unsetColor()
	if i.mode.IsDev() {
		i.log(colorGreen, "Reply", reply)
	}
}
