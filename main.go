package main

import (
	"github.com/kchunterdeluxe/algorithms/sort"
	"log"
	"math/rand"
	"time"
)

func randomInts(number int) []int {
	out := make([]int, number)

	for i := 0; i < number; i++ {
		out[i] = rand.Intn(100)
	}

	return out
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	input := randomInts(100000)

	start := time.Now()
	sort.Bubble(input)
	log.Printf("Time elapsed: %f seconds\n", time.Since(start).Seconds())
}
