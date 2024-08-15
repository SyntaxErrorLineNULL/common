package slices

import (
	"golang.org/x/exp/constraints"
	"sort"
)

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

// Contains checks if the provided element is present in the slice.
// It first sorts the slice and then performs a binary search to determine if the element exists.
// Returns true if the element is found, otherwise false.
func Contains[T constraints.Ordered](elements []T, element T) bool {
	// Check if the slice is nil. If it is, return false because there's nothing to search.
	if elements == nil {
		return false
	}

	// Sort the slice in ascending order.
	// Sorting is necessary for binary search to work correctly.
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] < elements[j]
	})

	// Use binary search to find the index of the element.
	// `sort.Search` will return the index of the first element greater than or equal to `element`.
	// If no such element is found, it returns the length of the slice.
	index := sort.Search(len(elements), func(i int) bool {
		return elements[i] >= element
	})

	// Validate the index to ensure it's within the bounds of the slice.
	// Check if the element at the found index matches the search element.
	// Return true if the element at the index equals the search element, otherwise false.
	return index < len(elements) && elements[index] == element
}

// Exclude removes all instances of a specified value from the provided slice.
// It creates a new slice containing only the elements that are not equal to the specified value.
// This approach efficiently constructs the result slice by reusing the original slice's underlying array,
// avoiding unnecessary memory allocations.
func Exclude[T comparable](elements []T, element T) []T {
	// Initialize the result slice with the same underlying array as the original slice.
	// This avoids unnecessary allocations and keeps the capacity the same.
	result := elements[:0]

	// Iterate over each item in the original slice.
	for _, item := range elements {
		// Check if the current item is not equal to the specified value to be excluded.
		if item != element {
			// Append the item to the result slice if it is not equal to the specified value.
			result = append(result, item)
		}
	}

	// Return the filtered slice with the specified value removed.
	return result
}
