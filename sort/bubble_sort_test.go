package sort

import (
	"reflect"
	gsort "sort"
	"testing"
)

func TestBubble(t *testing.T) {
	input := []int{2, 15, 3, 75, 18}
	output := Bubble(input)
	answer := make([]int, len(input))

	copy(answer, input)

	gsort.Ints(answer)

	if !reflect.DeepEqual(output, answer) {
		t.Errorf("Expected output == answer, got output: %v, answer: %v\n", output, answer)
	}
}
