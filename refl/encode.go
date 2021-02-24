package refl

// OBSOLETE

import (
	"reflect"
	"time"
)

func GetField(typ TLVVar, value reflect.Value) TLVField {
	t := value.Type()
	switch {
	// Special types
	case t == reflect.TypeOf([]byte(nil)):
		return NewBinary(typ, value.Bytes())
	case t == reflect.TypeOf(time.Millisecond):
		return NewNatural(typ, uint64(time.Duration(value.Int())/time.Millisecond))
	case t == reflect.TypeOf(Name(nil)):
		return NewNameField(typ, value.Interface().(Name))
	// Kind based check
	case t.Kind() == reflect.Uint64:
		return NewNatural(typ, value.Uint())
	case t.Kind() == reflect.Array || t.Kind() == reflect.Slice:
		return NewSequence(typ, value)
	case t.Kind() == reflect.Struct:
		return NewStruct(typ, value)
	// Pointer means optional
	case t.Kind() == reflect.Ptr:
		if value.Pointer() == uintptr(0) {
			return nil
		}
		return GetField(typ, value.Elem())
	default:
		panic("Unrecognized field")
	}
}

func DecodeField(buf []byte, t reflect.Type) (reflect.Value, error) {
	switch {
	case t == reflect.TypeOf([]byte(nil)):
		ret, err := ParseBinary(buf)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		return reflect.ValueOf(ret), nil
	case t == reflect.TypeOf(time.Millisecond):
		ret, err := ParseNatural(buf)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		return reflect.ValueOf(time.Duration(ret) * time.Millisecond), nil
	case t == reflect.TypeOf(Name(nil)):
		ret, err := ParseName(buf)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		return reflect.ValueOf(ret), nil
	case t.Kind() == reflect.Uint64:
		ret, err := ParseNatural(buf)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		return reflect.ValueOf(ret), nil
	case t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct:
		return ParseStruct(buf, t.Elem())
	case t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Struct:
		// The callee will append the return value to the slice, so no need to do anything
		return ParseStruct(buf, t.Elem())
	default:
		return reflect.ValueOf(nil), &DecodeError{Msg: "Unrecognized field"}
	}
}
