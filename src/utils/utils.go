package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Random(min, max int) (out int) {

	seed := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	out = seed.Intn(max-min) + min
	return
}

func Shuffle(in []byte) (out []byte) {

	for i := 0; i < len(in); i++ {
		rand := Random(0, len(in))
		in[i], in[rand] = in[rand], in[i]
	}
	out = in
	return
}

func Track(t time.Time) {

	startTime := t
	fmt.Printf("Job finished in %s\n", time.Now().Sub(startTime))
}
