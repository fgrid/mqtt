package mqtt_test

import (
	"bytes"
	"testing"

	"github.com/fgrid/mqtt"
)

func TestFixedHeaderWrite(t *testing.T) {
	fh := &mqtt.FixedHeader{
		PacketType:      0x01,
		Flags:           0x00,
		RemainingLength: 268435455,
	}

	var buf bytes.Buffer
	length, err := fh.Write(&buf)
	if err != nil {
		t.Errorf("could not write fixed header: %s", err.Error())
		return
	}
	if length != 5 {
		t.Errorf("fixed header encoding with unexpected length: %d (expected 5)", length)
		return
	}
	actual := buf.Bytes()
	expected := []byte{
		0x10,
		0xff, 0xff, 0xff, 0x7f,
	}
	if !bytes.Equal(actual, expected) {
		t.Errorf("unexpected fixed header encoding: % 2X", actual)
		t.Errorf("  expected fixed header encoding: % 2X", expected)
	}
}
