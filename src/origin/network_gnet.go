package origin

import "github.com/panjf2000/gnet"

type NetworkGnet struct {
	*gnet.EventServer
	Server * OriginServer
}

func (es *NetworkGnet) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {

	return
}

func (es *NetworkGnet) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	return
}
