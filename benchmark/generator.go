package benchmark

import (
	"fmt"
	"math/rand"
	"time"
)

type Case struct {
	Name    string
	Payload []byte
}

func Generator(cnt, payload, extraNameComponent int) (ret []Case) {
	extraName := ""
	for j := 0; j < extraNameComponent; j ++ {
		extraName += "a/"
	}

	ret = make([]Case, cnt)
	for i := 0; i < cnt; i++ {
		ret[i].Name = fmt.Sprintf("/go-ndn/test/%s%d", extraName, i+1)
		ret[i].Payload = make([]byte, payload)
		rand.Read(ret[i].Payload)
	}
	return
}

func Execute(cases []Case, executor func(Case)int) (time.Duration, int) {
	start := time.Now()
	totalBytes := 0
	for _, c := range cases {
		totalBytes += executor(c)
	}
	return time.Since(start), totalBytes
}
