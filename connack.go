package mqtt

import (
	"bytes"
	"fmt"
	"io"
)

type ConnACK struct {
	FixedHeader

	ConnectAcknowledgeFlags byte
	ConnectReturnCode       byte
}

func (c *ConnACK) WriteTo(w io.Writer) (int, error) {
	var payload bytes.Buffer
	Byte(c.ConnectAcknowledgeFlags).WriteTo(&payload)
	Byte(c.ConnectReturnCode).WriteTo(&payload)

	data := payload.Bytes()
	c.FixedHeader.RemainingLength = uint32(len(data))

	allCount := 0
	count, err := c.FixedHeader.WriteTo(w)
	if err != nil {
		return count, nil
	}
	allCount += count
	count, err = w.Write(data)
	if err != nil {
		return count + allCount, err
	}
	allCount += count
	return allCount, nil
}

func (c *ConnACK) SetPayload(raw []byte) error {
	if len(raw) != 2 {
		return fmt.Errorf("invalid payload for CONNACK (%d): % 2X", len(raw), raw)
	}
	c.ConnectAcknowledgeFlags = raw[0]
	c.ConnectReturnCode = raw[1]
	return nil
}
