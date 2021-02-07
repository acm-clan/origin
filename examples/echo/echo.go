package main

import (
	"origin/src/message"
	"origin/src/module"
	"origin/src/session"
)

type Echo struct {
	module.BaseModule
}

func (e *Echo) Echo(s *session.Session, msg * message.Message) {

}
