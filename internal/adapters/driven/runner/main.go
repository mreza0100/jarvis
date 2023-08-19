package runner

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/mreza0100/gptjarvis/internal/ports/runnerport"
	"github.com/mreza0100/gptjarvis/pkg/os"
)

type runner struct {
	conversationID string
	rootDir        string
}

func NewRunner() runnerport.Runner {
	conversationID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conversationID.String())
	return &runner{
		conversationID: conversationID.String(),
		rootDir:        "/home/mamzi/works/gpt-jarvis/tmp",
	}
}

func (r *runner) makeExecutable(path string) error {
	cmd := exec.Command("chmod", "+x", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (r *runner) writeScript(script string) (string, error) {
	time := getTime()

	pathName := path.Join(r.rootDir, r.conversationID, time)
	f, err := os.OpenFile(pathName, os.CreateMode)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	_, err = f.WriteString(script)
	return pathName, err
}

func (r *runner) getShebangLine(scriptPath string) (string, error) {
	// Read the first line of the script
	f, err := os.OpenFile(scriptPath, os.ReadMode)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		firstLine := scanner.Text()
		if strings.HasPrefix(firstLine, "#!") {
			return firstLine[2:], nil // Remove the "#!" prefix
		}
	}

	return "", fmt.Errorf("no shebang line found")
}

func (r *runner) RunScript(script string) (string, error) {
	path, err := r.writeScript(script)
	if err != nil {
		return "", err
	}
	err = r.makeExecutable(path)
	if err != nil {
		return "", err
	}

	shebangLine, err := r.getShebangLine(path)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("bash", "-c", shebangLine)
	cmd.Stdin = strings.NewReader(script) // Provide the script as input to the command
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr

	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	}
	return stdout.String(), nil
}
