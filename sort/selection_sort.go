package sort

// Selection sort is a sorting implementation that repeatedly steps through the list to be sorted,
// compares the other elements of the array for the minimum value, and swaps them. The Big-O performance statistics:
// Best case performance: O(n²)
// Average case performance: O(n²)
// Worst case performance O(n²)
func Selection(input []int) []int {

	for i := 0; i < len(input); i++ {
		minimum, swapIndex := input[i], i

		for j := (i + 1); j < len(input); j++ {
			if input[j] < minimum {
				minimum, swapIndex = input[j], j
			}
		}

		input[i], input[swapIndex] = minimum, input[i]
	}

	return input
}
