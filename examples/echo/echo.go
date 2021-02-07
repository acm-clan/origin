package main

import (
	"github.com/golang/protobuf/proto"
	"origin/examples/echo/pb"
	"origin/src/message"
	"origin/src/module"
	"origin/src/session"
)

type Echo struct {
	module.BaseModule
}

func (e *Echo) Echo(s *session.Session, msg * message.Message) {
	ec := &pb.Echo{}
	proto.Unmarshal(msg.RawData, ec)

	r := &pb.EchoReply{
		Msg: ec.Msg,
	}

	proto.Marshal(r)
}
