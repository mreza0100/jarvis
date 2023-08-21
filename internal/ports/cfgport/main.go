package cfgport

import "github.com/mreza0100/gptjarvis/internal/models"

type CfgProvider interface {
	GetCfg() *models.Configuration
}
