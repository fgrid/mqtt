package mqtt

import (
	"encoding/binary"
	"io"
)

type Integer uint16

func (i Integer) WriteTo(w io.Writer) (int, error) {
	raw := make([]byte, 2)
	binary.BigEndian.PutUint16(raw, uint16(i))
	return w.Write(raw)
}
