package main

import (
	"fmt"
	"time"

	"github.com/eric135/YaNFD/ndn"
	"github.com/eric135/YaNFD/ndn/security"
	"github.com/zjkmxy/ndn-encoding-test/benchmark"
	"github.com/zjkmxy/ndn-encoding-test/codegen"
	"github.com/zjkmxy/ndn-encoding-test/refl"
)

func main() {
	var cases []benchmark.Case

	fmt.Printf("#1: 1000000 x 100B\n")
	cases = benchmark.Generator(1000000, 100, 0)
	run(cases)

	fmt.Printf("#2: 45230 x 4000B\n")
	cases = benchmark.Generator(45230, 4000, 0)
	run(cases)

	fmt.Printf("#3: 1000000 x 1B, very long name\n")
	cases = benchmark.Generator(1000000, 1, 33)
	run(cases)

	return
}

func run(cases []benchmark.Case) {
	tim, totalBytes := benchmark.Execute(cases, blockEncode)
	fmt.Printf("block: \t\t%v\n", tim)
	tim, _ = benchmark.Execute(cases, reflEncode)
	fmt.Printf("reflection: \t%v\n", tim)
	tim, _ = benchmark.Execute(cases, codegenEncode)
	fmt.Printf("codegen: \t%v\n", tim)
	fmt.Printf("=== Total Bytes: %d ===\n", totalBytes)
	fmt.Println()
}

func blockEncode(c benchmark.Case) int {
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
	return len(wire)
}

func reflEncode(c benchmark.Case) int {
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
	return len(wire)
}

func codegenEncode(c benchmark.Case) int {
	data := &codegen.Data{
		Name: refl.NameFromStr(c.Name),
		MetaInfo: &codegen.MetaInfo{
			ContentType:     0,
			FreshnessPeriod: 4 * time.Second,
			FinalBlockID:    refl.NameComponentFromStr(8, "10000"),
		},
		Content: c.Payload,
		SignatureInfo: &codegen.SignatureInfo{
			SignatureType: 0,
		},
	}
	wire := data.Encode()
	// fmt.Printf("%+v\n", wire)
	return len(wire)
}
