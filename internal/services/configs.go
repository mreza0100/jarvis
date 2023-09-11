package services

import (
	"github.com/mreza0100/jarvis/internal/ports/srvport"
)

type ConfigSrv struct{}

func NewConfigSrv(req *srvport.OSServicesReq) srvport.ConfigService {
	return &ConfigSrv{}
}
