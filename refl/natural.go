package refl

import "encoding/binary"

// Natural represents a natural number field
type Natural struct {
	typ    TLVVar
	value  uint64
	length uint
}

func NewNatural(typ TLVVar, value uint64) Natural {
	var l uint
	switch {
	case value <= 0xff:
		l = 1
	case value <= 0xffff:
		l = 2
	case value < 0xffffffff:
		l = 4
	default:
		l = 8
	}
	return Natural{
		typ:    typ,
		value:  value,
		length: l,
	}
}

func (v Natural) Length() uint {
	return v.typ.Length() + 1 + v.length
}

func (v Natural) Encode(buf []byte) uint {
	pos := v.typ.Encode(buf)
	buf[pos] = byte(v.length)
	pos += 1
	switch v.length {
	case 1:
		buf[pos] = byte(v.value)
	case 2:
		binary.BigEndian.PutUint16(buf[pos:pos+2], uint16(v.value))
	case 4:
		binary.BigEndian.PutUint32(buf[pos:pos+4], uint32(v.value))
	case 8:
		binary.BigEndian.PutUint64(buf[pos:pos+8], uint64(v.value))
	}
	return pos + v.length
}
