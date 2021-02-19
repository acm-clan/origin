package session

const (
	PacketHeaderSize = 5
)

type BodyMetadata struct {
	MessageID uint32

}

type Packet struct {
	Type byte
	Code byte
	Size uint32
	Body []byte
}

func ReadPacket(stream *ReadStream) * Packet {
	return nil
}
