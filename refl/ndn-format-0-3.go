package refl

import (
	"crypto/sha256"
	"reflect"
	"time"
)

type MetaInfo struct {
	ContentType     uint64        `tlv:"0x18"`
	FreshnessPeriod time.Duration `tlv:"0x19"`
	FinalBlockID    NameComponent `tlv:"0x1a"`
}

type SignatureInfo struct {
	SignatureType uint64 `tlv:"0x1b"`
}

type Data struct {
	Name           Name           `tlv:"0x07"`
	MetaInfo       *MetaInfo      `tlv:"0x14"`
	Content        []byte         `tlv:"0x15"`
	SignatureInfo  *SignatureInfo `tlv:"0x16"`
	SignatureValue []byte         `tlv:"0x17"`
}

func (v *Data) Encode() []byte {
	// Just quick and dirty hack
	v.SignatureInfo.SignatureType = 0
	sha := sha256.New()
	v.SignatureValue = make([]byte, sha.Size())
	stru := NewStruct(TLVVar(0x06), reflect.ValueOf(v).Elem())
	l := stru.Length()
	ret := make([]byte, l)
	stru.Encode(ret)
	pos := 1
	switch ret[pos] {
	case 0xfd:
		pos += 3
	case 0xfe:
		pos += 5
	case 0xff:
		pos += 9
	default:
		pos += 1
	}
	end := int(l) - sha.Size() - 2
	sha.Write(ret[pos:end])
	sha.Sum(ret[0 : end+2])
	return ret
}

func DecodeData(buf []byte) *Data {
	typ, size := DecodeTLVVar(buf)
	if typ != 0x06 {
		return nil
	}
	pos := size
	l, size := DecodeTLVVar(buf[pos:])
	pos += size
	if pos+uint(l) != uint(len(buf)) {
		return nil
	}
	ret, err := ParseStruct(buf[pos:pos+uint(l)], reflect.TypeOf(Data{}))
	if err != nil {
		return nil
	}
	return (ret.Interface()).(*Data)
}
