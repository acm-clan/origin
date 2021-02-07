package origin

import (
	"origin/src/module"
)

type OriginServer struct {
	Option * Option
}

func Default() *OriginServer {
	or := &OriginServer{
		Option: NewOption(),
	}
	return or
}

func (or *OriginServer) Run() {
	
	println("origin server")
}

func (or *OriginServer) Handler(h module.Module) {

}
