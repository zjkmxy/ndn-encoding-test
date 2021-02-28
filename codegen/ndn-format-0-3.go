package codegen

import (
	"crypto/sha256"
	"time"

	"github.com/zjkmxy/ndn-encoding-test/refl"
)

type MetaInfo struct {
	//+field:natural
	ContentType uint64 `tlv:"0x18"`
	//+field:time
	FreshnessPeriod time.Duration `tlv:"0x19"`
	//+field:binary
	FinalBlockID refl.NameComponent `tlv:"0x1a"`
}

type SignatureInfo struct {
	//+field:natural
	SignatureType uint64 `tlv:"0x1b"`
}

type Data struct {
	//+field:name
	Name refl.Name `tlv:"0x07"`
	//+field:struct:MetaInfo
	MetaInfo *MetaInfo `tlv:"0x14"`
	//+field:binary
	Content []byte `tlv:"0x15"`
	//+field:struct:SignatureInfo
	SignatureInfo *SignatureInfo `tlv:"0x16"`
	//+field:signature
	SignatureValue []byte `tlv:"0x17"`
}

func (v *Data) Encode() []byte {
	e := DataEncoder{}
	e.init(v)
	l := e.length
	tlvLen := refl.TLVVar(l)
	ret := make([]byte, 1+tlvLen.Length()+l)
	pos := refl.TLVVar(0x06).Encode(ret)
	pos += tlvLen.Encode(ret[pos:])

	sha := sha256.New()
	end := pos + e.encodeInto(v, ret[pos:]) - uint(sha.Size()) - 2
	sha.Write(ret[pos:end])
	sha.Sum(ret[0 : end+2])
	return ret
}

func DecodeData(buf []byte) *Data {
	typ, size := refl.DecodeTLVVar(buf)
	if typ != 0x06 {
		return nil
	}
	pos := size
	l, size := refl.DecodeTLVVar(buf[pos:])
	pos += size
	if pos+uint(l) != uint(len(buf)) {
		return nil
	}
	ret, l2 := ParseData(buf[pos:], false)
	if uint(l) != l2 {
		return nil
	}
	return ret
}
