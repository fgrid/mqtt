package mqtt

import (
	"bytes"
	"io"
)

type Connect struct {
	FixedHeader

	ProtocolName string

	ProtocolLevel byte

	UserNameFlag bool
	PasswordFlag bool
	WillRetain   bool
	WillQoS      byte
	WillFlag     bool
	CleanSession bool

	ClientIdentifier string
	WillTopic        string
	WillMessage      string
	UserName         string
	Password         string

	KeepAlive uint16
}

func (c *Connect) WriteTo(w io.Writer) (int, error) {
	var payload bytes.Buffer
	String(c.ProtocolName).WriteTo(&payload)
	Byte(c.ProtocolLevel).WriteTo(&payload)
	Byte(0x00).WriteTo(&payload)
	String(c.ClientIdentifier).WriteTo(&payload)
	Integer(c.KeepAlive).WriteTo(&payload)

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
