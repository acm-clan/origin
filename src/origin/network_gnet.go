package origin

import (
	"github.com/panjf2000/gnet"
	"origin/src/logger"
	"origin/src/session"
	"time"
)

type NetworkGnet struct {
	*gnet.EventServer
	Server * OriginServer
}

func (es *NetworkGnet) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	logger.Infof("Client Open:%v", c.RemoteAddr())
	// create session
	s := session.CreateSession(c)
	c.SetContext(s)
	return
}

func (es *NetworkGnet) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	logger.Infof("React:%v", string(frame))
	s := session.Get(c)
	s.OnNewData(frame)
	return
}

// 服务器启动完成
func (es *NetworkGnet) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	logger.Debugf("OnInitComplete")
	return
}

func (es *NetworkGnet) OnShutdown(svr gnet.Server) {
	logger.Debugf("OnShutdown")
}

func (es *NetworkGnet) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	logger.Debugf("OnClosed")
	return
}

func (es *NetworkGnet) PreWrite() {
	logger.Debugf("PreWrite")
}

func (es *NetworkGnet) Tick() (delay time.Duration, action gnet.Action) {
	logger.Debugf("Tick %v", delay)
	return
}