package services

import (
	"github.com/mreza0100/jarvis/internal/ports/srvport"
	"github.com/mreza0100/jarvis/internal/services/boot"
)

func NewServices(req *srvport.ServicesReq) *srvport.Services {
	return &srvport.Services{
		BootService: boot.NewBootSrv(req),
	}
}
