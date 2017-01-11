package mqtt

import (
	"encoding/binary"
	"io"
)

type String string

func (s String) WriteTo(w io.Writer) (int, error) {
	dest := make([]byte, 2+len(s))
	binary.BigEndian.PutUint16(dest, uint16(len(s)))
	copy(dest[2:], []byte(s))
	return w.Write(dest)
}
