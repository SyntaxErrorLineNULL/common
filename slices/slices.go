package slices

// Merge concatenates two slices into a single slice.
// It creates a new slice with a length equal to the sum of the lengths of the input slices.
// The function copies all elements from the first slice followed by all elements from the second slice into the new slice,
// and returns this combined slice.
func Merge[T any](first, second []T) []T {
	// Allocate a new slice with enough capacity to hold all elements from both input slices.
	list := make([]T, len(first)+len(second))
	// Copy all elements from the first slice into the new slice.
	copy(list, first)
	// Copy all elements from the second slice into the new slice, starting right after the first slice's elements.
	copy(list[len(first):], second)
	// Return the combined slice containing elements from both input slices.
	return list
}
