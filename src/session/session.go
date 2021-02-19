package session

import (
	"bytes"
	"github.com/panjf2000/gnet"
	"origin/src/logger"
)

type Session struct {
	Buf bytes.Buffer
	Connection gnet.Conn
}

func bytesToInt(b []byte) uint32 {
	result := 0
	for _, v := range b {
		result = result<<8 + int(v)
	}
	return uint32(result)
}

func CreateSession(conn gnet.Conn)*Session{
	s := &Session{
		Buf: bytes.Buffer{},
		Connection: conn,
	}
	return s
}

func Get(conn gnet.Conn) *Session{
	c := conn.Context()
	return c.(*Session)
}

func (s* Session) PeekPacket() *Packet {
	if s.Buf.Len() < PacketHeaderSize {
		return nil
	}

	//read header
	b := make([]byte, PacketHeaderSize)
	s.Buf.Read(b)

	size := bytesToInt(b[2:])

	if uint32(s.Buf.Len()) < size {
		s.Buf.UnreadRune()
		return nil
	}

	data := make([]byte, size)
	s.Buf.Read(data)

	p := &Packet{
		Type: b[0],
		Code: b[1],
		Size: size,
		Body: data,
	}

	return p
}

func (s* Session) OnNewData(data []byte){
	s.Buf.Write(data)

	for p := s.PeekPacket(); p != nil; p = s.PeekPacket() {
		s.OnNewPacket(p)
	}
}

func (s* Session) OnNewPacket(p * Packet){
	logger.Debugf("new packet:%v", p.Size)
}



