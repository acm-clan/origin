package origin

type OriginServer struct {
}

func Default() *OriginServer {
	return &OriginServer{}
}

func (or *OriginServer) Run() {
	println("origin server")
}
