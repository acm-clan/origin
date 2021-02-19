package client

import (
	"net"
	"sync"
)

type Message struct {
	data string
}

// BenchClient send message to server
type BenchClient struct {
	ClientID          int64
	MessageCount      int64
	MessageWriteCount int64
	MessageReadCount  int64
	WaitGroup         *sync.WaitGroup
	MessageSize       int64
	chWrite           chan *Message
	chClose           chan struct{}
	Connection        net.Conn
	Sync              bool
}

func NewBenchClient(CID int64, wg *sync.WaitGroup, messageCount int64, messageSize int64) *BenchClient {
	return &BenchClient{
		ClientID:          CID,
		MessageCount:      messageCount,
		MessageReadCount:  0,
		MessageWriteCount: 0,
		MessageSize:       messageSize,
		WaitGroup:         wg,
		chWrite:           make(chan *Message),
		chClose:           make(chan struct{}),
		Sync:              true,
	}
}

// Start start a bench client
func (bc *BenchClient) Start(ip string, port int) {
	bc.StartSync(ip, port)
}
