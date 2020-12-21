package refl

import (
	"reflect"
	"strconv"
)

// Struct represents a structure field or a TLV model field
type Struct struct {
	typ    TLVVar
	length uint
	values []TLVField
}

func NewStruct(typ TLVVar, value reflect.Value) *Struct {
	v := value
	t := v.Type()
	length := uint(0)
	numFields := v.NumField()
	fields := make([]TLVField, numFields)
	for i := 0; i < numFields; i++ {
		f := t.Field(i)
		typStr, exists := f.Tag.Lookup("tlv")
		if !exists {
			fields[i] = nil
			continue
		}
		typVal, err := strconv.ParseUint(typStr, 0, 0)
		if err != nil {
			panic(err)
		}

		fields[i] = GetField(TLVVar(typVal), v.Field(i))
		if fields[i] != nil {
			length += fields[i].Length()
		}
	}

	return &Struct{
		typ:    typ,
		length: length,
		values: fields,
	}
}

func (v *Struct) Length() uint {
	if v.typ > 0 {
		return v.typ.Length() + TLVVar(v.length).Length() + v.length
	} else {
		return v.length
	}
}

func (v *Struct) Encode(buf []byte) uint {
	pos := uint(0)
	if v.typ > 0 {
		pos += v.typ.Encode(buf)
		pos += TLVVar(v.length).Encode(buf[pos:])
	}
	for _, f := range v.values {
		if f != nil {
			pos += f.Encode(buf[pos:])
		}
	}
	return pos
}
