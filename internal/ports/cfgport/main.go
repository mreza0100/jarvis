package cfgport

import "github.com/mreza0100/jarvis/internal/models"

type CfgProvider interface {
	GetCfg() *models.Configuration
}
