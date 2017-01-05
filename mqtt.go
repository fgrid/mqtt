package mqtt

import "encoding/binary"

func PutInteger(dest []byte, data uint16) {
	binary.BigEndian.PutUint16(dest, data)
}

func Integer(src []byte) uint16 {
	return binary.BigEndian.Uint16(src)
}

func PutString(dest []byte, data string) {
	binary.BigEndian.PutUint16(dest, uint16(len(data)))
	copy(dest[2:], []byte(data))
}

func String(src []byte) string {
	length := binary.BigEndian.Uint16(src[:2])
	return string(src[2 : 2+length])
}
