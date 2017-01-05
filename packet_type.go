package mqtt

// PacketType is represented as a 4-bit unsigned value.
type PacketType byte

const (
	Reserved = iota
	CONNECT
	CONNACK
	PUBLISH
	PUBACK
	PUBREC
	PUBREL
	PUBCOMP
	SUBSCRIBE
	SUBACK
	UNSUBSCRIBE
	UNSUBACK
	PINGREQ
	PINGRESP
	DISCONNECT
	Reserved2
)

var packetType = [...]string{
	"Reserved",
	"CONNECT",
	"CONNACK",
	"PUBLISH",
	"PUBACK",
	"PUBREC",
	"PUBREL",
	"PUBCOMP",
	"SUBSCRIBE",
	"SUBACK",
	"UNSUBSCRIBE",
	"UNSUBACK",
	"PINGREQ",
	"PINGRESP",
	"DISCONNECT",
	"Reserved2",
}

// String representation of the given packet type
func (pt PacketType) String() string {
	if int(pt) >= len(packetType) {
		return "!!! undefined !!!"
	}
	return packetType[pt]
}
