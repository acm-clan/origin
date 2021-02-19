package client

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"origin/examples/echo/pb"
	"origin/src/logger"
)

// Decode packet data length byte to int(Big end)
func bytesToInt(b []byte) int {
	result := 0
	for _, v := range b {
		result = result<<8 + int(v)
	}
	return result
}

// Encode packet data length to bytes(Big end)
func intToBytes(n int) []byte {
	buf := make([]byte, 3)
	buf[0] = byte((n >> 16) & 0xFF)
	buf[1] = byte((n >> 8) & 0xFF)
	buf[2] = byte(n & 0xFF)
	return buf
}

func encodeMessage(msg string) []byte {
	r := &pb.Echo{
		Msg: msg,
	}
	w := &bytes.Buffer{}
	data, err := proto.Marshal(r)
	if err != nil {
		panic("proto marshal failed")
	}
	//type data
	w.WriteByte(0x02)
	//code
	w.WriteByte(0x01)
	//length
	w.Write(intToBytes(len(data)))
	//body
	w.Write(data)
	return w.Bytes()
}

func (bc *BenchClient) benchSync(ip string, port int) {
	connection, err := net.Dial("tcp", ip+":"+fmt.Sprint(port))
	if err != nil {
		logger.Error("[bench] error connect server: ", err)
		return
	}

	bc.Connection = connection

	response := bufio.NewReader(connection)

	for i := 0; i < int(bc.MessageCount); i++ {
		msg := "hello world\n"
		logger.Debug("[bench] msg: ", msg)

		_, err := bc.Connection.Write(encodeMessage(msg))
		if err != nil {
			logger.Errorf("write error: %v", err)
			break
		}

		ret, err := response.ReadBytes(byte('\n'))

		logger.Debugf("[echo] %v %v ", i+1, string(ret))

		if err != nil {
			logger.Errorf("read error: %v", err)
			break
		}
	}

}

// Start start a bench client
func (bc *BenchClient) StartSync(ip string, port int) {
	logger.Debugf("[client] start %v", bc.ClientID)
	bc.benchSync(ip, port)
	if bc.WaitGroup != nil {
		bc.WaitGroup.Done()
	}
}
