package refl

// TLVField is an encoder for a field in s structure
type TLVField interface {
	Length() uint
	Encode([]byte) uint
}

type DecodeError struct {
	Msg string
}

func (e *DecodeError) Error() string {
	return e.Msg
}
