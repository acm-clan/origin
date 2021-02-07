package origin

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"origin/src/logger"
	"origin/src/module"
)



type OriginServer struct {
	Option * Option
	Modules []module.Module
}

func Default() *OriginServer {
	or := &OriginServer{
		Option: NewOption(),
	}
	logger.InitLogger("debug")
	return or
}

func (or *OriginServer) Run() {
	switch or.Option.NetworkType {
	case NetworkTypeGNet:
		s := new(NetworkGnet)
		addr := fmt.Sprintf("tcp://:%d", or.Option.ServerPort)

		logger.Infof("Origin server is serving with port %d", or.Option.ServerPort)
		err := gnet.Serve(s, addr, gnet.WithMulticore(true))

		if err != nil {

		}
	case NetworkTypeGoNet:
		logger.Info("not implement yet")
	}

}

func (or *OriginServer) Module(mod module.Module) {
	or.Modules = append(or.Modules, mod)
}
