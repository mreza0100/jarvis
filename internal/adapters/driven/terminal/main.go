package terminal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/sashabaranov/go-openai"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/pkg/tools"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/terminalport"
)

type terminal struct {
	mode models.Mode

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

type TerminalReq struct {
	CfgProvider cfgport.CfgProvider

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewTerminal(req TerminalReq) terminalport.Terminal {
	return &terminal{
		mode: req.CfgProvider.GetConfigs().Mode,

		stdin:  req.Stdin,
		stdout: req.Stdout,
		stderr: req.Stderr,
	}
}

func (i *terminal) GetUserInput() (string, error) {
	fmt.Print("\n==> ")
	reader := bufio.NewReader(i.stdin)
	userPrompt, _ := reader.ReadString('\n')
	return strings.TrimSpace(userPrompt), nil
}

func (i *terminal) setColor(color color) {
	fmt.Print(color)
}

func (i *terminal) unsetColor() {
	fmt.Printf("%s[%dm", "\x1b", 0)
}

func (i *terminal) log(color color, title string, message any) {
	i.setColor(color)
	defer i.unsetColor()

	fmt.Printf("%s\n%+v\n", title, message)
}

func (i *terminal) Error(err error) {
	i.log(colorRed, "\n", err.Error())
}

func (i *terminal) printInsights(rateLimitInsights *openai.RateLimitHeaders) {
	i.setColor(colorPurple)
	fmt.Printf("\n")
	fmt.Printf("Limit Tokens: %v", rateLimitInsights.LimitTokens)
	fmt.Printf(" - Remaining Tokens: %v", rateLimitInsights.RemainingTokens)
	fmt.Printf(" - Reset Tokens in: %v", rateLimitInsights.ResetTokens)
	fmt.Printf("\n")
	fmt.Printf("Limit Requests: %v", rateLimitInsights.LimitRequests)
	fmt.Printf(" - Remaining Requests: %v", rateLimitInsights.RemainingRequests)
	fmt.Printf(" - Reset Requests in: %v", rateLimitInsights.ResetRequests)
	fmt.Printf("\n")
	i.unsetColor()
}

func (i *terminal) PrintReply(message string, rateLimitInsights *openai.RateLimitHeaders) {
	i.printInsights(rateLimitInsights)
	width, _, err := tools.GetTerminalSize()
	if err != nil {
		log.Fatal(err)
		return
	}
	renderedMarkdown := markdown.Render(message, width, 0)

	i.log(colorCyan, "\n", string(renderedMarkdown))
}

func (i *terminal) Script(script interface{}) {
	i.setColor(colorRed)
	defer i.unsetColor()
	if i.mode.IsDev() {
		fmt.Printf("\n- Script:\n")
		fmt.Printf("script.String()=%+v\n", script)
		fmt.Printf("\n--\n")
	}
}

func (i *terminal) ScriptResults(result interface{}) {
	if i.mode.IsDev() {
		jsonResult, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
		}
		i.log(colorBlue, "Script Results", string(jsonResult))
	}
}

func (i *terminal) Reply(reply interface{}) {
	i.setColor(colorPurple)
	defer i.unsetColor()
	if i.mode.IsDev() {
		i.log(colorGreen, "Reply", reply)
	}
}
