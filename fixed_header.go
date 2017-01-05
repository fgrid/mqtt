package mqtt

import "io"

// The FixedHeader is present in all MQTT Control Packets.
type FixedHeader struct {
	// PacketType represented as a 4-bit unsigned value and declares the packet type
	PacketType byte

	// Flags are specific to each MQTT Control Packet type.
	Flags byte

	// RemainingLength is the number of bytes remaining within the currenct packet, including data
	// in the variable header and the payload. The Remaining Length does not include the bytes used to
	// encode the Remaining Length.
	//
	// The Remaining Length is encoded using a variable length encoding scheme which uses a single byte
	// for values up to 127. Larger values are handled as follows. The least significant seven bits of
	// each byte encode the data, and the most significant bit is used to indicate that there are
	// following bytes in the representation. Thus each byte encodes 128 values and a "continuation bit".
	// The maximum number of bytes in the Remaining Length field is four.
	RemainingLength uint32
}

// Write the fixed header to the given writer using the MQTT wire protocol
func (fh *FixedHeader) Write(w io.Writer) (n int, err error) {

	n, err = w.Write([]byte{(fh.PacketType << 4) | fh.Flags})
	if err != nil {
		return
	}
	var second int
	second, err = w.Write(fh.encodeLength())
	n += second
	return
}

func (fh *FixedHeader) encodeLength() []byte {
	dest := make([]byte, 0, 4)
	rl := fh.RemainingLength
	for more := true; more; {
		encodedByte := byte(rl % 128)
		rl = rl / 128
		if rl > 0 {
			encodedByte |= 128
		} else {
			more = false
		}
		dest = append(dest, encodedByte)
	}
	return dest
}
