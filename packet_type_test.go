package mqtt_test

import (
	"fmt"

	"github.com/fgrid/mqtt"
)

func ExamplePacketType_String() {
	for i := 0; i <= 17; i++ {
		pt := mqtt.PacketType(i)
		fmt.Printf("%02d - %s\n", pt, pt)
	}
	// Output:
	// 00 - Reserved
	// 01 - CONNECT
	// 02 - CONNACK
	// 03 - PUBLISH
	// 04 - PUBACK
	// 05 - PUBREC
	// 06 - PUBREL
	// 07 - PUBCOMP
	// 08 - SUBSCRIBE
	// 09 - SUBACK
	// 10 - UNSUBSCRIBE
	// 11 - UNSUBACK
	// 12 - PINGREQ
	// 13 - PINGRESP
	// 14 - DISCONNECT
	// 15 - Reserved2
	// 16 - !!! undefined !!!
	// 17 - !!! undefined !!!
}
