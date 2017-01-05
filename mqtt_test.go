package mqtt_test

import (
	"bytes"
	"testing"

	"github.com/fgrid/mqtt"
)

func TestPutInteger(t *testing.T) {
	dest := make([]byte, 2)
	mqtt.PutInteger(dest, 15)
	expected := []byte{0x00, 0x0f}
	if !bytes.Equal(dest, expected) {
		t.Errorf("unexpected integer encoding: % 2X", dest)
		t.Errorf("  expected integer encoding: % 2X", expected)
	}
}

func TestInteger(t *testing.T) {
	src := []byte{0x00, 0x0f}
	actual := mqtt.Integer(src)
	if actual != 15 {
		t.Errorf("unexpected integer after decoding: %d", actual)
	}
}

const exampleString = "A𪛔" // 1.5.3.1

func TestPutString(t *testing.T) {
	dest := make([]byte, 7)
	mqtt.PutString(dest, exampleString)
	expected := []byte{
		0x00, 0x05,
		0x41, 0xf0, 0xaa, 0x9b, 0x94,
	}
	if !bytes.Equal(dest, expected) {
		t.Errorf("unexpected string encoding: % 2X", dest)
		t.Errorf("  expected string encoding: % 2X", expected)
	}
}

func TestString(t *testing.T) {
	src := []byte{
		0x00, 0x05,
		0x41, 0xf0, 0xaa, 0x9b, 0x94,
	}
	actual := mqtt.String(src)
	if actual != exampleString {
		t.Errorf("unexpecte string after decoding: %s", actual)
	}
}
