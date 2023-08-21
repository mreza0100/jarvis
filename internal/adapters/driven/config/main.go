package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"path"
	"strings"

	"github.com/mreza0100/gptjarvis/internal/models"
	"github.com/mreza0100/gptjarvis/internal/ports/cfgport"
	"github.com/mreza0100/gptjarvis/pkg/os"
)

var loadedConfigs *models.Configuration

const (
	configDirName  = ".jarvis"
	configFileName = "config.json"
	historyDirName = "history"
)

type configs struct {
	cfg *models.Configuration
}

func NewConfigProvider() cfgport.CfgProvider {
	if loadedConfigs != nil {
		return &configs{cfg: loadedConfigs}
	}

	cfg := &configs{cfg: &models.Configuration{}}
	cfg.setEnvConfigs()
	cfg.setSavedConfigs()
	cfg.setConstantConfigs()
	return cfg
}

func (c *configs) GetCfg() *models.Configuration {
	return c.cfg
}

func (c *configs) setEnvConfigs() {
	c.cfg.Mode = getModeFromEnv("MODE")
}

func (c *configs) setConstantConfigs() {
	c.cfg.ConstantConfigs = &models.ConstantConfigs{
		RootDirPath:    configDirName,
		HistoryDirName: historyDirName,
	}
}

func (c *configs) setSavedConfigs() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jarvisDir := path.Join(homeDir, configDirName)

	savedConfigs := new(models.SavedConfigs)

	cfgF, err := os.OpenFile(path.Join(jarvisDir, configFileName), os.ReadCreateMode)
	if err != nil {
		log.Fatal(err)
	}

	rawContent, err := io.ReadAll(cfgF)
	if err != nil {
		log.Fatal(err)
	}
	content := string(rawContent)
	if content == "" {
		// TODO: Marshal savedConfigs and write it instead of {}
		content = "{}"
		cfgF, err := os.OpenFile(path.Join(jarvisDir, configFileName), os.AppendMode)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := io.WriteString(cfgF, content); err != nil {
			log.Fatal(err)
		}
	}
	if err := json.Unmarshal([]byte(content), savedConfigs); err != nil {
		log.Fatal(err)
	}
}

func getModeFromEnv(key string) models.Mode {
	value := strings.ToLower(os.Getenv(key))
	switch value {
	case "dev":
		return models.Modes.Dev
	default:
		return models.Modes.Prod
	}
}
