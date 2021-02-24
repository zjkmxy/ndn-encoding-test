package refl

// Binary represents a binary string field
type Binary struct {
	typ   TLVVar
	value []byte
}

func NewBinary(typ TLVVar, value []byte) Binary {
	return Binary{
		typ:   typ,
		value: value,
	}
}

func (v Binary) Length() uint {
	return v.typ.Length() + TLVVar(len(v.value)).Length() + uint(len(v.value))
}

func (v Binary) Encode(buf []byte) uint {
	pos := v.typ.Encode(buf)
	l := uint(len(v.value))
	pos += TLVVar(l).Encode(buf[pos:])
	copy(buf[pos:pos+l], v.value)
	return pos + l
}

func ParseBinary(buf []byte) ([]byte, error) {
	return buf, nil
}
