package refl

import (
	"reflect"
)

// Sequence represents an array field
type Sequence struct {
	typ    TLVVar
	length uint
	values []TLVField
}

func NewSequence(typ TLVVar, value reflect.Value) Sequence {
	l := value.Len()
	values := make([]TLVField, l)
	length := uint(0)
	for i := 0; i < l; i++ {
		values[i] = GetField(typ, value.Index(i))
		if values[i] != nil {
			length += values[i].Length()
		}
	}
	return Sequence{
		typ:    typ,
		length: length,
		values: values,
	}
}

func (v Sequence) Length() uint {
	return v.length
}

func (v Sequence) Encode(buf []byte) uint {
	pos := uint(0)
	for _, f := range v.values {
		if f != nil {
			pos += f.Encode(buf[pos:])
		}
	}
	return pos
}
