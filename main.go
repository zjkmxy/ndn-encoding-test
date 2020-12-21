package main

import (
	"fmt"
	"time"

	"github.com/eric135/YaNFD/ndn"
	"github.com/eric135/YaNFD/ndn/security"
	"github.com/zjkmxy/ndn-encoding-test/benchmark"
	"github.com/zjkmxy/ndn-encoding-test/refl"
)

func main() {
	cases := benchmark.Generator(1000000, 100)
	// cases := benchmark.Generator(1, 1)
	tim := benchmark.Execute(cases, blockEncode)
	fmt.Printf("%v\n", tim)
	tim = benchmark.Execute(cases, reflEncode)
	fmt.Printf("%v\n", tim)

	return
}

func blockEncode(c benchmark.Case) {
	name, err := ndn.NameFromString(c.Name)
	if err != nil {
		panic(err)
	}
	data := ndn.NewData(name, c.Payload)
	metaInfo := ndn.NewMetaInfo()
	metaInfo.SetFinalBlockID(ndn.NewGenericNameComponent([]byte("10000")))
	metaInfo.SetFreshnessPeriod(4 * time.Second)
	metaInfo.SetContentType(0)
	data.SetMetaInfo(metaInfo)
	data.SetSignatureInfo(ndn.NewSignatureInfo(security.DigestSha256Type))
	block, err := data.Encode()
	if err != nil {
		panic(err)
	}
	wire, err := block.Wire()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", wire)
	noop(wire)
}

func reflEncode(c benchmark.Case) {
	data := &refl.Data{
		Name: refl.NameFromStr(c.Name),
		MetaInfo: &refl.MetaInfo{
			ContentType:     0,
			FreshnessPeriod: 4 * time.Second,
			FinalBlockID:    refl.NameComponentFromStr(8, "10000"),
		},
		Content: c.Payload,
		SignatureInfo: &refl.SignatureInfo{
			SignatureType: 0,
		},
	}
	wire := data.Encode()
	// fmt.Printf("%+v\n", wire)
	noop(wire)
}

func noop(v interface{}) {
}
