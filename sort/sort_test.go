package sort

import (
	"math/rand"
	"reflect"
	gsort "sort"
	"testing"
)

type Test struct {
	input  []int
	output []int
}

var tests []Test

func init() {
	// Test input of random integers (10, 100, 1000 elements)
	ten := make([]int, 10)
	hundred := make([]int, 100)
	thousand := make([]int, 1000)

	for i := 0; i < 10; i++ {
		ten[i] = rand.Intn(10)
	}

	for i := 0; i < 100; i++ {
		hundred[i] = rand.Intn(100)
	}

	for i := 0; i < 1000; i++ {
		thousand[i] = rand.Intn(1000)
	}

	sortedTen := make([]int, 10)
	sortedHundred := make([]int, 100)
	sortedThousand := make([]int, 1000)

	copy(sortedTen, ten)
	copy(sortedHundred, hundred)
	copy(sortedThousand, thousand)

	gsort.Ints(sortedTen)
	gsort.Ints(sortedHundred)
	gsort.Ints(sortedThousand)

	tests = []Test{
		Test{
			input:  ten,
			output: sortedTen,
		},
		Test{
			input:  hundred,
			output: sortedHundred,
		},
		Test{
			input:  thousand,
			output: sortedThousand,
		},
	}
}

func TestBubble(t *testing.T) {

	for _, test := range tests {
		output := Bubble(test.input)

		if reflect.DeepEqual(output, test.output) == false {
			t.Errorf("Expected %v, got %v", test.output, output)
		}
	}

}

func TestSelection(t *testing.T) {

	for _, test := range tests {
		output := Selection(test.input)

		if reflect.DeepEqual(output, test.output) == false {
			t.Errorf("Expected %v, got %v", test.output, output)
		}
	}

}
