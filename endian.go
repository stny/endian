// Copyright (c) 2014 Naoya SATO

package endian

var MiddleEndian middleEndian

type middleEndian struct{}

func (middleEndian) Uint16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1]) << 8
}

func (middleEndian) PutUint16(b []byte, v uint16) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
}

func (middleEndian) Uint32(b []byte) uint32 {
	return uint32(b[2]) | uint32(b[3]) << 8 | uint32(b[0]) << 16 | uint32(b[1]) << 24
}

func (middleEndian) PutUint32(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 24)
	b[2] = byte(v)
	b[3] = byte(v >> 8)
}

func (middleEndian) Uint64(b []byte) uint64 {
	return uint64(b[6]) | uint64(b[7]) << 8 | uint64(b[4]) << 16 | uint64(b[5]) << 24 |
	uint64(b[2]) << 32 | uint64(b[3]) << 40 | uint64(b[0]) << 48 | uint64(b[1]) << 56
}

func (middleEndian) PutUint64(b []byte, v uint64) {
	b[0] = byte(v >> 48)
	b[1] = byte(v >> 56)
	b[2] = byte(v >> 32)
	b[3] = byte(v >> 40)
	b[4] = byte(v >> 16)
	b[5] = byte(v >> 24)
	b[6] = byte(v)
	b[7] = byte(v >> 8)
}

func (middleEndian) String() string { return "MiddleEndian" }
