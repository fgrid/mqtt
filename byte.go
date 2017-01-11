package mqtt

import "io"

type Byte byte

func (b Byte) WriteTo(w io.Writer) (int, error) {
	raw := make([]byte, 1)
	raw[0] = byte(b)
	return w.Write(raw)
}
