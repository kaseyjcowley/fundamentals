package sort

// Bubble sort is a naive sorting implementation that repeatedly steps through the list to be sorted,
// compares each pair of adjacent items and swaps them if they are in the wrong order. The Big-O performance statistics:
// Best case performance: O(n)
// Average case performance: O(n²)
// Worst case performance O(n²)
func Bubble(input []int) []int {

	for {
		swapped := false

		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return input

}
