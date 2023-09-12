package cfgport

import "github.com/mreza0100/jarvis/internal/models"

type CfgProvider interface {
	RefreshCfg(path *string) *models.Configuration
	GetCfg() *models.Configuration
}
