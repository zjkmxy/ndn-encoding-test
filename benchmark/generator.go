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

func Generator(cnt, payload int) (ret []Case) {
	ret = make([]Case, cnt)
	for i := 0; i < cnt; i++ {
		ret[i].Name = fmt.Sprintf("/go-ndn/test/%d", i+1)
		ret[i].Payload = make([]byte, payload)
		rand.Read(ret[i].Payload)
	}
	return
}

func Execute(cases []Case, executor func(Case)) time.Duration {
	start := time.Now()
	for _, c := range cases {
		executor(c)
	}
	return time.Since(start)
}
