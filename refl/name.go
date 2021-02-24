package refl

import (
	"strings"
)

type NameComponent = []byte

type Name []NameComponent

func NameComponentFromStr(typ uint, value string) NameComponent {
	l := uint(len(value))
	buf := make([]byte, TLVVar(typ).Length()+TLVVar(l).Length()+l)
	pos := TLVVar(typ).Encode(buf)
	pos += TLVVar(l).Encode(buf[pos:])
	copy(buf[pos:pos+l], value)
	return buf
}

func NameFromStr(name string) Name {
	// This is just quick and dirty hack. Conventions are not supported
	slices := strings.Split(name, "/")
	ret := make([]NameComponent, len(slices))
	for i, s := range slices[1:] {
		ret[i] = NameComponentFromStr(8, s)
	}
	return ret
}

type NameField struct {
	typ    TLVVar
	value  Name
	length uint
}

func NewNameField(typ TLVVar, value Name) NameField {
	l := 0
	for _, c := range value {
		l += len(c)
	}
	return NameField{
		typ:    typ,
		value:  value,
		length: uint(l),
	}
}

func (v NameField) Length() uint {
	return v.typ.Length() + TLVVar(v.length).Length() + v.length
}

func (v NameField) Encode(buf []byte) uint {
	pos := v.typ.Encode(buf)
	pos += TLVVar(v.length).Encode(buf[pos:])
	for _, c := range v.value {
		copy(buf[pos:], c)
		pos += uint(len(c))
	}
	return pos
}

func ParseName(buf []byte) (Name, error) {
	pos := 0
	ret := make([]NameComponent, 0)
	for pos < len(buf) {
		_, size1 := DecodeTLVVar(buf[pos:])
		l, size2 := DecodeTLVVar(buf[pos+int(size1):])
		totalLen := int(size1) + int(size2) + int(l)
		component := buf[pos : pos+totalLen]
		ret = append(ret, component)
		pos += totalLen
	}
	return ret, nil
}
