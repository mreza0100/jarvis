package cfgport

import "github.com/mreza0100/jarvis/internal/models"

type CfgProvider interface {
	LoadConfig() *models.Configuration
	GetConfigs() *models.Configuration
	LoadStoredFile(fileName string) (string, error)
	BootstrapHostConfig(trunk bool) error
}
type CfgService interface {
	Bootstrap(trunk bool) error
}
