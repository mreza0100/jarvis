package config

import (
	"encoding/json"
	"io"
	"log"
	"path"
	"strings"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/pkg/os"
	promptstore "github.com/mreza0100/jarvis/promptstore"
)

var configFileEmptySchehma = models.ConfigFile{
	PostgresConfig: models.PostgresConfig{
		PostgresConnConfig: models.PostgresConnConfig{
			Host:     "",
			Port:     5432,
			Username: "",
			Password: "",
			Database: "",
		},
	},
}

const (
	configDirName  = ".jarvis"
	configFileName = "config.json"
	historyDirName = "history"
)

type configProvider struct {
	cfg            *models.Configuration
	configFilePath *string
}

func NewConfigProvider(path *string) cfgport.CfgProvider {
	cfg := &configProvider{
		configFilePath: path,
		cfg: &models.Configuration{
			ConfigFile: &models.ConfigFile{},
		},
	}
	cfg.RefreshCfg(path)
	return cfg
}

func (c *configProvider) LoadSavedFile(fileName string) (string, error) {
	content, err := promptstore.ModelsFS.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	templ := string(content)
	return templ, nil
}

func (c *configProvider) RefreshCfg(path *string) *models.Configuration {
	if path != nil {
		c.loadConfigFile(*path)
	}
	c.setEnvConfigs()
	c.setConstantConfigs()
	return c.cfg
}

func (c *configProvider) GetConfigs() *models.Configuration {
	return c.cfg
}

func (c *configProvider) setEnvConfigs() {
	c.cfg.Token = os.Getenv("OPEN_API_KEY")
	c.cfg.Mode = getModeFromEnv("MODE")
}

func (c *configProvider) setConstantConfigs() {
	c.cfg.HistoryDirName = historyDirName
	c.cfg.RootDirName = configDirName
}

func (c *configProvider) loadConfigFile(p string) {
	cfgF, err := os.OpenFile(path.Join(p), os.ReadCreateMode)
	if err != nil {
		log.Fatal(err)
	}

	rawContent, err := io.ReadAll(cfgF)
	if err != nil {
		log.Fatal(err)
	}
	if len(rawContent) == 0 {
		jsonConfigSchema, err := json.Marshal(configFileEmptySchehma)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.OpenFile(p, os.AppendMode)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := io.WriteString(f, string(jsonConfigSchema)); err != nil {
			log.Fatal(err)
		}
		rawContent = jsonConfigSchema
	}
	if err := json.Unmarshal(rawContent, c.cfg.ConfigFile); err != nil {
		log.Fatal(err)
	}
}

func getModeFromEnv(key string) models.Mode {
	switch value := strings.ToLower(os.Getenv(key)); value {
	case "dev":
		return models.Modes.Dev
	default:
		return models.Modes.Prod
	}
}

// func (c *configs) setSavedConfigs() {
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	jarvisDir := path.Join(homeDir, configDirName)

// 	savedConfigs := new(models.SavedConfigs)

// cfgF, err := os.OpenFile(path.Join(jarvisDir, configFileName), os.ReadCreateMode)
// if err != nil {
// 	log.Fatal(err)
// }

// rawContent, err := io.ReadAll(cfgF)
// if err != nil {
// 	log.Fatal(err)
// }
// content := string(rawContent)
// if content == "" {
// 	// TODO: Marshal savedConfigs and write it instead of {}
// 	content = "{}"
// 	cfgF, err := os.OpenFile(path.Join(jarvisDir, configFileName), os.AppendMode)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if _, err := io.WriteString(cfgF, content); err != nil {
// 		log.Fatal(err)
// 	}
// }
// if err := json.Unmarshal([]byte(content), savedConfigs); err != nil {
// 	log.Fatal(err)
// }
// }
