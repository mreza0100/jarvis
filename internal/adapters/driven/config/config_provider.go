package config

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"path"
	"strings"

	"github.com/mreza0100/jarvis/internal/models"
	"github.com/mreza0100/jarvis/internal/ports/cfgport"
	"github.com/mreza0100/jarvis/pkg/os"
	"github.com/mreza0100/jarvis/store/preprompts"
)

var EmptyConfigFileSchema = &models.ConfigFile{
	Postgres: &models.PostgresConfig{
		Config: &models.ChatConfig{Model: "gpt-4", Temperature: 1},
		PostgresConnConfig: &models.PostgresConnConfig{
			Host:     "",
			Port:     5432,
			Username: "",
			Password: "",
			Database: "",
		},
	},
	OS: &models.OSConfig{
		Config: &models.ChatConfig{Model: "gpt-4", Temperature: 1},
	},
}

const (
	configDirName  = ".jarvis"
	configFileName = ".jarvisrc.json"
	historyDirName = "history"
)

type configProvider struct {
	cfg             *models.Configuration
	configFilePath  string
	hasConfigLoaded bool
}

func NewConfigProvider() cfgport.CfgProvider {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	cfg := &configProvider{
		configFilePath: path.Join(homePath, configDirName, configFileName),
		cfg: &models.Configuration{
			ConfigFile: &models.ConfigFile{},
		},
	}
	return cfg
}

func (c *configProvider) LoadStoredFile(fileName string) (string, error) {
	content, err := preprompts.ModelsFS.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	templ := string(content)
	return templ, nil
}

func (c *configProvider) BootstrapHostConfig(trunk bool) error {
	if cfgFileExists := os.FileExists(c.configFilePath); !trunk && cfgFileExists {
		// TODO: sortout and intgrate all the errors
		return errors.New("A configuration file already exists. To replace the existing configuration with a new one, please rerun the command using the '--trunk' flag. This action will erase the current configuration and create a new one.")
	}

	jsonConfigSchema, err := json.MarshalIndent(EmptyConfigFileSchema, "", "	")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(c.configFilePath, os.TruncMode)
	if err != nil {
		return err
	}
	if _, err := io.WriteString(f, string(jsonConfigSchema)); err != nil {
		return err
	}

	return nil
}

func (c *configProvider) LoadConfig() *models.Configuration {
	c.loadConfigFile()
	c.setEnvConfigs()
	c.setConstantConfigs()

	c.hasConfigLoaded = true

	return c.cfg
}

func (c *configProvider) GetConfigs() *models.Configuration {
	if !c.hasConfigLoaded {
		c.LoadConfig()
	}
	return c.cfg
}

func (c *configProvider) setEnvConfigs() {
	c.cfg.Mode = getModeFromEnv("MODE")
}

func (c *configProvider) setConstantConfigs() {
	c.cfg.HistoryDirName = historyDirName
	c.cfg.RootDirName = configDirName
}

func (c *configProvider) loadConfigFile() {
	cfgF, err := os.OpenFile(c.configFilePath, os.ReadCreateMode)
	if err != nil {
		log.Fatal(err)
	}

	rawContent, err := io.ReadAll(cfgF)
	if err != nil {
		log.Fatal(err)
	}
	if len(rawContent) == 0 {
		log.Fatal("Config file does not exist or invalid. please bootstrap the config")
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

// const configFileName = "config.json"
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
