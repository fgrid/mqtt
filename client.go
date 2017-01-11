package mqtt

import (
	"log"
	"net"
	"time"
)

type Client struct {
	ID        string
	KeepAlive time.Duration
}

var DefaultClient = &Client{}

func (c *Client) Connect(network, address string) error {
	con, err := net.Dial(network, address)
	if err != nil {
		return err
	}
	connect := Connect{
		FixedHeader: FixedHeader{
			PacketType: CONNECT,
		},
		ProtocolName:     "MQTT",
		ProtocolLevel:    0x04,
		ClientIdentifier: c.ID,
		KeepAlive:        uint16(c.KeepAlive.Seconds()),
	}
	_, err = connect.WriteTo(con)

	var fh FixedHeader
	if err := fh.ReadFrom(con); err != nil {
		return err
	}
	payload := make([]byte, fh.RemainingLength)
	ca := ConnACK{FixedHeader: fh}
	if err := ca.SetPayload(payload); err != nil {
		return err
	}
	log.Printf("connect returncode: % 2X", ca.ConnectReturnCode)
	return nil

}
