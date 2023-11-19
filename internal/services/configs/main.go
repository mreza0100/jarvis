package configs

import "github.com/mreza0100/jarvis/internal/ports/cfgport"

type Configuration struct {
	configProvider cfgport.CfgProvider
}

func NewConfigController(cfgProvider cfgport.CfgProvider) cfgport.CfgService {
	return &Configuration{
		configProvider: cfgProvider,
	}
}

func (c *Configuration) Bootstrap(trunk bool) error {
	return c.configProvider.BootstrapHostConfig(trunk)
}
