package utils

import (
	"math/rand"
	"time"
)

func Random(min, max int) (out int) {

	seed := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	out = seed.Intn(max-min) + min
	return
}

func RandomGo(min, max int, ch chan int) {

	for {
		seed := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		ch <- seed.Intn(max-min) + min
	}
}

func getRandomGo(N int, ch chan int) (out int) {

	for i := 0; i < N; i++ {
		out = <-ch
	}
	return
}

func Shuffle(in []byte) []byte {

	out := make([]byte, len(in))
	copy(out, in[:])
	for i := 0; i < len(out); i++ {
		rand := Random(0, len(out))
		out[i], out[rand] = out[rand], out[i]
	}
	return out
}
