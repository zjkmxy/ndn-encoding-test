package main

import (
	"fmt"
	"time"

	"github.com/eric135/YaNFD/ndn"
	"github.com/eric135/YaNFD/ndn/security"
	"github.com/eric135/YaNFD/ndn/tlv"
	"github.com/zjkmxy/ndn-encoding-test/benchmark"
	"github.com/zjkmxy/ndn-encoding-test/codegen"
	"github.com/zjkmxy/ndn-encoding-test/refl"
)

const (
	testEncoding = true
	testDecoding = true
)

func main() {
	var cases []benchmark.Case

	if testEncoding {
		fmt.Printf("========== ENCODING ==========\n")
		fmt.Printf("#1: 1000000 x 100B\n")
		cases = benchmark.GenerateEncodingCases(1000000, 100, 0)
		run(cases)

		fmt.Printf("#2: 1000000 x 4000B\n")
		cases = benchmark.GenerateEncodingCases(1000000, 4000, 0)
		run(cases)

		fmt.Printf("#3: 1000000 x 10B, very long name\n")
		cases = benchmark.GenerateEncodingCases(1000000, 10, 30)
		run(cases)
	}

	if testDecoding {
		fmt.Printf("\n========== DECODING ==========\n")
		fmt.Printf("#1: 1000000 x 100B\n")
		cases = benchmark.GenerateDecodingCases(1000000, 100, 0)
		runDecoding(cases)

		fmt.Printf("#2: 1000000 x 4000B\n")
		cases = benchmark.GenerateDecodingCases(1000000, 4000, 0)
		runDecoding(cases)

		fmt.Printf("#3: 1000000 x 10B, very long name\n")
		cases = benchmark.GenerateDecodingCases(1000000, 10, 30)
		runDecoding(cases)
	}
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

func runDecoding(cases []benchmark.Case) {
	tim, totalBytes := benchmark.Execute(cases, blockDecode)
	fmt.Printf("block: \t\t%v\n", tim)
	tim, _ = benchmark.Execute(cases, reflDecode)
	fmt.Printf("reflection: \t%v\n", tim)
	tim, _ = benchmark.Execute(cases, codegenDecode)
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

func blockDecode(c benchmark.Case) int {
	wire, size, err := tlv.DecodeBlock(c.Payload)
	if err != nil {
		panic(err)
	}
	data, err := ndn.DecodeData(wire, false)
	if data == nil || err != nil {
		panic(err)
	}
	return int(size)
}

func codegenDecode(c benchmark.Case) int {
	data := codegen.DecodeData(c.Payload)
	if data == nil {
		panic("Unable to parse data")
	}
	return len(c.Payload)
}

func reflDecode(c benchmark.Case) int {
	data := refl.DecodeData(c.Payload)
	if data == nil {
		panic("Unable to parse data")
	}
	return len(c.Payload)
}
