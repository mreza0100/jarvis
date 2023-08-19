package config

import (
	"github.com/mreza0100/gptjarvis/internal/ports/cfgport"
	cfgPort "github.com/mreza0100/gptjarvis/internal/ports/configs"
)

type configs struct {
	cfg *cfgPort.Configs
}

func NewConfigProvider() cfgport.CfgProvider {
	return &configs{}
}

func (c *configs) SetCfg(cfg *cfgPort.Configs) {
	c.cfg = cfg
}

func (c *configs) GetCfg() *cfgPort.Configs {
	return c.cfg
}
