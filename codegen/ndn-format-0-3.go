package codegen

import (
	"time"

	"github.com/zjkmxy/ndn-encoding-test/refl"
)

type MetaInfo struct {
	ContentType     uint64             `tlv:"0x18"`
	FreshnessPeriod time.Duration      `tlv:"0x19"`
	FinalBlockID    refl.NameComponent `tlv:"0x1a"`
}

type SignatureInfo struct {
	SignatureType uint64 `tlv:"0x1b"`
}

type Data struct {
	Name           refl.Name      `tlv:"0x07"`
	MetaInfo       *MetaInfo      `tlv:"0x14"`
	Content        []byte         `tlv:"0x15"`
	SignatureInfo  *SignatureInfo `tlv:"0x16"`
	SignatureValue []byte         `tlv:"0x17"`
}

func (v *Data) Encode() []byte {
	e := DataEncoder{}
	e.init(v)
	ret := make([]byte, e.length)
	e.encodeInto(v, ret)
	return ret
}
