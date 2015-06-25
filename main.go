package main

import (
	"fmt"
	"github.com/kchunterdeluxe/algorithms/sort"
)

func main() {
	input := []int{5, 10, 3, 17, 9, 34}

	output := sort.Bubble(input)

	fmt.Printf("%v", output)
}
