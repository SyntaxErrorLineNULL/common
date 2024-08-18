package test

// CreateSequenceWithRepeats generates a slice of integers with a specified size.
// The slice contains a repeated element at every 100th position, while other positions
// are filled with their respective indices.
func CreateSequenceWithRepeats(size, repeatedElement int) []int {
	// Initialize a slice with the specified size.
	slice := make([]int, size)

	// Iterate over each index in the slice.
	for i := 0; i < size; i++ {
		// If the index is a multiple of 100, insert the repeated element.
		if i%100 == 0 {
			slice[i] = repeatedElement
		} else {
			// Otherwise, insert the index value itself.
			slice[i] = i
		}
	}

	// Return the generated slice.
	return slice
}

// CreateSequenceWithoutRepeats generates a slice of integers with a specified size,
// ensuring that no element is repeated at positions that are multiples of 100.
func CreateSequenceWithoutRepeats(size int) []int {
	// Initialize an empty slice with a predefined capacity.
	slice := make([]int, 0, size)

	// Iterate over each index up to the specified size.
	for i := 0; i < size; i++ {
		// Only include indices that are not multiples of 100.
		if i%100 != 0 {
			// Append the index to the slice.
			slice = append(slice, i)
		}
	}

	// Return the generated slice.
	return slice
}
