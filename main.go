package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/named-data/YaNFD/ndn"
	"github.com/named-data/YaNFD/ndn/security"
	"github.com/named-data/YaNFD/ndn/tlv"
	enc "github.com/zjkmxy/go-ndn/pkg/encoding"
	gondn "github.com/zjkmxy/go-ndn/pkg/ndn"
	"github.com/zjkmxy/go-ndn/pkg/ndn/spec_2022"
	gondnsec "github.com/zjkmxy/go-ndn/pkg/security"
	"github.com/zjkmxy/go-ndn/pkg/utils"
	"github.com/zjkmxy/ndn-encoding-test/benchmark"
	"github.com/zjkmxy/ndn-encoding-test/codegen"
	"github.com/zjkmxy/ndn-encoding-test/refl"
)

const (
	testEncoding   = true
	testDecoding   = true
	signOnEncoding = false
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	var cases []benchmark.Case

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if testEncoding {
		fmt.Printf("========== ENCODING ==========\n")
		fmt.Printf("#1: 1000000 x 100B\n")
		cases = benchmark.GenerateEncodingCases(1000000, 100, 0, signOnEncoding)
		run(cases)

		fmt.Printf("#2: 1000000 x 4000B\n")
		cases = benchmark.GenerateEncodingCases(1000000, 4000, 0, signOnEncoding)
		run(cases)

		fmt.Printf("#3: 1000000 x 10B, very long name\n")
		cases = benchmark.GenerateEncodingCases(1000000, 10, 30, signOnEncoding)
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
	tim, _ = benchmark.Execute(cases, gondnEncode)
	fmt.Printf("go-ndn: \t%v\n", tim)
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
	tim, _ = benchmark.Execute(cases, gondnDecode)
	fmt.Printf("go-ndn: \t%v\n", tim)
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
	if c.IfSign {
		data.SetSignatureInfo(ndn.NewSignatureInfo(security.DigestSha256Type))
	} else {
		data.SetSignatureInfo(ndn.NewSignatureInfo(security.SignatureNullType))
	}
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
	wire := data.Encode(c.IfSign)
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
	if c.IfSign {
		data.SignatureValue = []byte{}
	} else {
		data.SignatureInfo.SignatureType = 200
	}
	wire := data.Encode()
	// fmt.Printf("%+v\n", wire)
	return len(wire)
}

func gondnEncode(c benchmark.Case) int {
	spec := spec_2022.Spec{}
	name, _ := enc.NameFromStr(c.Name)
	signer := gondn.Signer(nil)
	if c.IfSign {
		signer = gondnsec.NewSha256Signer()
	}
	wire, _, _ := spec.MakeData(
		name,
		&gondn.DataConfig{
			ContentType: utils.IdPtr(gondn.ContentTypeBlob),
			Freshness:   utils.IdPtr(4 * time.Second),
			FinalBlockID: &enc.Component{
				Typ: 8,
				Val: []byte("10000"),
			},
		},
		enc.Wire{c.Payload},
		signer,
	)
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

func gondnDecode(c benchmark.Case) int {
	spec := spec_2022.Spec{}
	data, _, _ := spec.ReadData(enc.NewBufferReader(c.Payload))
	if data == nil {
		panic("Unable to parse data")
	}
	return len(c.Payload)
}
