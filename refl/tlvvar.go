package refl

import "encoding/binary"

// TLVVar represents a TLV variable (type/length)
type TLVVar uint

func (v TLVVar) Length() uint {
	switch {
	case v <= 0xfc:
		return 1
	case v <= 0xffff:
		return 3
	case v < 0xffffffff:
		return 5
	default:
		return 9
	}
}

func (v TLVVar) Encode(buf []byte) uint {
	switch {
	case v <= 0xfc:
		buf[0] = byte(v)
		return 1
	case v <= 0xffff:
		buf[0] = 0xfd
		binary.BigEndian.PutUint16(buf[1:], uint16(v))
		return 3
	case v < 0xffffffff:
		buf[0] = 0xfe
		binary.BigEndian.PutUint32(buf[1:], uint32(v))
		return 5
	default:
		buf[0] = 0xff
		binary.BigEndian.PutUint64(buf[1:], uint64(v))
		return 9
	}
}

func DecodeTLVVar(buf []byte) (TLVVar, uint) {
	v := buf[0]
	switch v {
	case 0xfd:
		return TLVVar(binary.BigEndian.Uint16(buf[1:3])), 3
	case 0xfe:
		return TLVVar(binary.BigEndian.Uint32(buf[1:5])), 5
	case 0xff:
		return TLVVar(binary.BigEndian.Uint64(buf[1:9])), 9
	default:
		return TLVVar(v), 1
	}
}
