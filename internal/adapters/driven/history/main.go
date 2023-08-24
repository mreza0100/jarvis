package history

import (
	"fmt"
	"io"
	"log"
	"path"
	"time"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/internal/ports/historyport"
	"github.com/mreza0100/jarvis/pkg/os"
)

type history struct {
	mode             models.Mode
	historyDirPath   string
	chatIndexTracker int
}

func NewHistory(cfgProvider cfgport.CfgProvider) historyport.History {
	constCfg := cfgProvider.GetCfg().ConstantConfigs
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	startTime := time.Now().Format("2006-01-02_15-04-05")
	historyDirPath := path.Join(homeDir, constCfg.RootDirPath, constCfg.HistoryDirName, startTime)
	return &history{
		mode:             cfgProvider.GetCfg().Mode,
		historyDirPath:   historyDirPath,
		chatIndexTracker: 0,
	}
}

func (h *history) SavePrompt(prompt *models.Prompt) {
	h.saveRecord(&models.HistoryRecord{
		Prompt: prompt,
	})
}

func (h *history) SaveResponse(response *models.Response) {
	h.saveRecord(&models.HistoryRecord{
		Response: response,
	})
}

func (h *history) saveRecord(record *models.HistoryRecord) {
	f, err := os.OpenFile(path.Join(h.historyDirPath, fmt.Sprintf("%v", h.chatIndexTracker)), os.CreateMode)
	if err != nil {
		log.Fatal(err)
	}
	rawRecord, err := record.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(f, rawRecord)
	if err != nil {
		log.Fatal(err)
	}
	h.chatIndexTracker++
}
