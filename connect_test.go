package mqtt_test

import (
	"bytes"
	"testing"

	"github.com/fgrid/mqtt"
)

func TestCONNECTWrite(t *testing.T) {
	c := mqtt.Connect{
		FixedHeader: mqtt.FixedHeader{
			PacketType: mqtt.CONNECT,
		},
		ProtocolName:     "MQTT",
		ProtocolLevel:    0x04,
		UserNameFlag:     true,
		PasswordFlag:     true,
		WillFlag:         true,
		CleanSession:     true,
		KeepAlive:        10,
		ClientIdentifier: "myClientID",
		WillTopic:        "*",
		UserName:         "myUserName",
		Password:         "myPassword",
	}
	var buf bytes.Buffer
	length, err := c.WriteTo(&buf)
	if err != nil {
		t.Errorf("could not write connect packet: %s", err.Error())
		return
	}
	if length != 24 {
		t.Errorf("could not write whole CONNECT packet: written %d bytes - expected 20", length)
		t.Errorf("actual: % 2X", buf.Bytes())
		return
	}

}
