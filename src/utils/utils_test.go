package utils

import (
	"testing"
	"time"
)

func TestRandom(t *testing.T) {

	var testSet = []struct {
		min int
		max int
	}{
		{1, 20},
		{2, 200},
		{1, 20},
		{1, 100},
		{3, 5},
		{41, 54},
		{-11, 42},
		{0, 20},
	}

	for _, tt := range testSet {
		res := Random(tt.min, tt.max)
		if res >= tt.max || res < tt.min {
			t.Log("out of range! min: %d, max: %d", tt.min, tt.max)
			t.Fail()
		}
	}

}

func TestShuffle(t *testing.T) {

	in := []byte("askdsakdiuwqqSDGGEWFDS%#@<>dji")
	out := Shuffle(in)
	if len(out) != len(in) {
		t.Log("length not match\n")
		t.Fail()
	}

	//test smililarity of two shuffled strings
	var similar = make(map[byte]int)
	count := 0
	for i := 0; i < len(in); i++ {
		if in[i] == out[i] {
			similar[in[i]]++
			count++
			t.Log("same position byte and times:", string(in[i]), ",", similar[in[i]], "\n")
		}
	}
	q := float32(count) / float32(len(in))
	if q >= 0.4 {
		t.Log("too similar!", q*100, "%")
	}
}

func BenchmarkRandom(b *testing.B) {

	b.N = 999999
	for i := 0; i < b.N; i++ {
		out := Random(1, 100)
		b.Log(out)
	}

}

func BenchmarkRandomGo(b *testing.B) {
	ch := make(chan int)
	b.N = 999999
	go RandomGo(1, 100, ch)
	go getRandomGo(b.N, ch)
	time.Sleep(1e9)

}
