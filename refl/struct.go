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
		typVal := extractTlvTag(f)
		if typVal == 0 {
			fields[i] = nil
			continue
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

func isSlice(fieldType reflect.Type) bool {
	return fieldType.Kind() == reflect.Slice &&
		fieldType.Elem().Kind() == reflect.Struct
}

func ParseStruct(buf []byte, structType reflect.Type) (reflect.Value, error) {
	fieldOffset := make(map[uint64]int)
	for i := 0; i < structType.NumField(); i++ {
		f := structType.Field(i)
		typVal := extractTlvTag(f)
		if typVal != 0 {
			fieldOffset[typVal] = i
		}
	}

	progress := -1
	bufLen := uint(len(buf))
	ret := reflect.New(structType)
	var l TLVVar
	for pos := uint(0); pos < bufLen; pos += uint(l) {
		typ, size := DecodeTLVVar(buf[pos:])
		pos += size
		l, size = DecodeTLVVar(buf[pos:])
		pos += size
		// NEW
		fieldIdx, exists := fieldOffset[uint64(typ)]
		if exists && fieldIdx >= progress {
			f := structType.Field(fieldIdx)
			val, err := DecodeField(buf[pos:pos+uint(l)], f.Type)
			if err != nil {
				return reflect.ValueOf(nil), err
			}

			if fieldIdx > progress {
				progress = fieldIdx
				if !isSlice(f.Type) {
					ret.Elem().Field(fieldIdx).Set(val)
				} else {
					slice := reflect.MakeSlice(f.Type.Elem(), 0, 1)
					// TODO: Should we de-ptr val first?
					slice = reflect.Append(slice, val)
					ret.Elem().Field(fieldIdx).Set(slice)
				}
			} else if isSlice(f.Type) {
				progress = fieldIdx
				slice := ret.Elem().Field(fieldIdx)
				slice = reflect.Append(slice, val)
				ret.Elem().Field(fieldIdx).Set(slice)
			} else {
				// Duplicated
				exists = false
			}
		} else {
			// Out-of-order
			exists = false
		}
		if !exists {
			// Error handling
		}
	}
	// Return the pointer to ret
	return ret, nil
}

func extractTlvTag(f reflect.StructField) uint64 {
	typStr, exists := f.Tag.Lookup("tlv")
	if !exists {
		return 0
	}
	typVal, err := strconv.ParseUint(typStr, 0, 0)
	if err != nil {
		panic(err)
	}
	return typVal
}
