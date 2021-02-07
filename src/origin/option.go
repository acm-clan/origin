package origin

const (
	ProtocolTypeProtobuf = iota
	ProtocolTypeJson
)

const (
	NetworkTypeGNet = iota
	NetworkTypeGoNet
)

type Option struct {
	ProtocolType int8
	ServerPort int32
	NetworkType int8
}

func NewOption()*Option{
	return &Option{
		ProtocolType: ProtocolTypeProtobuf,
		ServerPort: 9600,
		NetworkType: NetworkTypeGNet,
	}
}