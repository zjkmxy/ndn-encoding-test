package benchmark

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zjkmxy/ndn-encoding-test/refl"
)

type Case struct {
	Name    string
	Payload []byte
}

func GenerateEncodingCases(cnt, payload, extraNameComponent int) (ret []Case) {
	extraName := ""
	for j := 0; j < extraNameComponent; j++ {
		extraName += "a/"
	}

	ret = make([]Case, cnt)
	for i := 0; i < cnt; i++ {
		ret[i].Name = fmt.Sprintf("/go-ndn/test/%s%010d", extraName, i+1)
		ret[i].Payload = make([]byte, payload)
		rand.Read(ret[i].Payload)
	}
	return
}

func Execute(cases []Case, executor func(Case) int) (time.Duration, int) {
	start := time.Now()
	totalBytes := 0
	for _, c := range cases {
		totalBytes += executor(c)
	}
	return time.Since(start), totalBytes
}

func GenerateDecodingCases(cnt, payload, extraNameComponent int) (ret []Case) {
	var namePartOne = [...]byte{8, 6, 103, 111, 45, 110, 100, 110, 8, 4, 116, 101, 115, 116}
	var nameFlexiblePart = [...]byte{8, 1, 97}
	// For simplicity, here I use fixed number
	var namePartTwo = [...]byte{8, 10, 48, 48, 48, 48, 48, 48, 48, 48, 48, 49}
	var beforeContent = [...]byte{20, 16, 24, 1, 0, 25, 2, 15, 160, 26, 7, 8, 5, 49, 48, 48, 48, 48}
	var signature = [...]byte{22, 3, 27, 1, 0, 23, 32,
		163, 60, 194, 104, 248, 118, 79, 231, 250, 185, 231, 57, 251, 224, 71, 220,
		255, 220, 10, 60, 77, 181, 242, 75, 60, 83, 208, 210, 38, 129, 234, 198}

	nameLength := refl.TLVVar(len(namePartOne) + len(namePartTwo) + len(nameFlexiblePart)*extraNameComponent)
	contentLength := refl.TLVVar(payload)
	dataLength := refl.TLVVar(1 + nameLength.Length() + uint(nameLength) + uint(len(beforeContent)) +
		1 + contentLength.Length() + uint(contentLength) + uint(len(signature)))
	totalLength := 1 + dataLength.Length() + uint(dataLength)

	buf := make([]byte, totalLength)
	// Data
	buf[0] = 0x06
	pos := uint(1)
	pos += dataLength.Encode(buf[pos:])
	// Name
	buf[pos] = 0x07
	pos++
	pos += nameLength.Encode(buf[pos:])
	copy(buf[pos:], namePartOne[:])
	pos += uint(len(namePartOne))
	for j := 0; j < extraNameComponent; j++ {
		copy(buf[pos:], nameFlexiblePart[:])
		pos += uint(len(nameFlexiblePart))
	}
	copy(buf[pos:], namePartTwo[:])
	pos += uint(len(namePartTwo))
	// Before content
	copy(buf[pos:], beforeContent[:])
	pos += uint(len(beforeContent))
	// Content
	buf[pos] = 0x15
	pos++
	pos += contentLength.Encode(buf[pos:])
	rand.Read(buf[pos : pos+uint(payload)])
	pos += uint(payload)
	// Signature
	copy(buf[pos:], signature[:])
	pos += uint(len(signature))

	ret = make([]Case, cnt)
	for i := 0; i < cnt; i++ {
		ret[i].Payload = make([]byte, totalLength)
		copy(ret[i].Payload, buf)
	}
	return
}
