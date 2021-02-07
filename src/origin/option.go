package origin

const (
	ProtocolTypeProtobuf = iota
	ProtocolTypeJson
)

type Option struct {
	ProtocolType int8
	ServerPort int32
}

func NewOption()*Option{
	return &Option{
		ProtocolType: ProtocolTypeProtobuf,
		ServerPort: 9600,
	}
}