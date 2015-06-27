package main

import (
	"fmt"
	"github.com/kchunterdeluxe/fundamentals/sort"
	"math/rand"
	"time"
)

func main() {
	items := 100000
	input := rand.Perm(items)

	// Bubble Sort
	start := time.Now()

	fmt.Printf("Starting bubble sort on %d items...\n", items)
	sort.Bubble(input)
	fmt.Printf("Bubble sort took %s to complete on %d items.\n\n", time.Since(start), items)

	// Selection Sort
	start = time.Now()

	fmt.Printf("Starting selection sort on %d items...\n", items)
	sort.Selection(input)
	fmt.Printf("Selection sort took %s to complete on %d items.\n\n", time.Since(start), items)
}
