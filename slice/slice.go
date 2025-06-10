package core

import (
	"sort"

	"golang.org/x/exp/constraints"
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

// Contains checks if the provided element is present in the slice.
// It first sorts the slice and then performs a binary search to determine if the element exists.
// Returns true if the element is found, otherwise false.
func Contains[T constraints.Ordered](elements []T, element T) bool {
	// Check if the slice is nil. If it is, return false because there's nothing to search.
	if elements == nil {
		return false
	}

	// Create a copy of the input slice to avoid modifying the original slice.
	copiedElements := make([]T, len(elements))
	copy(copiedElements, elements)

	// Sort the slice in ascending order.
	// Sorting is necessary for binary search to work correctly.
	sort.Slice(copiedElements, func(i, j int) bool {
		return copiedElements[i] < copiedElements[j]
	})

	// Use binary search to find the index of the element.
	// `sort.Search` will return the index of the first element greater than or equal to `element`.
	// If no such element is found, it returns the length of the slice.
	index := sort.Search(len(copiedElements), func(i int) bool {
		return copiedElements[i] >= element
	})

	// Validate the index to ensure it's within the bounds of the slice.
	// Check if the element at the found index matches the search element.
	// Return true if the element at the index equals the search element, otherwise false.
	return index < len(copiedElements) && copiedElements[index] == element
}

// Map applies a transformation function to each element of a slice and returns a new slice with the transformed elements.
// This function takes a slice of type A and applies the provided transformation function to each element,
// resulting in a new slice of type B containing the transformed values.
// It ensures that the order of elements is preserved in the output slice, and all transformations are applied correctly.
func Map[A, B any](elements []A, fn func(A) B) []B {
	// Create a slice of type B with the same length as the input slice.
	// This ensures the result slice has sufficient capacity to store the transformed elements.
	result := make([]B, len(elements))

	// Iterate over the input slice to transform each element.
	// The index and value of each element are retrieved during the iteration.
	for i, v := range elements {
		// Apply the transformation function to the current element and store the result in the corresponding position in the result slice.
		// This preserves the original order of the elements while applying the transformation.
		result[i] = fn(v)
	}

	// Return the resulting slice containing the transformed elements.
	return result
}

// Filter filters a slice of elements based on a provided predicate function.
// It iterates over each element in the input slice and applies the predicate function to it.
// If the predicate returns true for an element, that element is included in the result.
// The result is a new slice containing only the elements that satisfied the predicate function.
// This function is generic and works with any type of slice and predicate.
func Filter[T any](elements []T, fn func(T) bool) []T {
	var result []T

	// Iterate over each element in the input slice.
	// For each element, apply the predicate function.
	// If the predicate returns true, append the element to the result slice.
	for _, v := range elements {
		// Check if the current element satisfies the predicate function fn.
		// The predicate function fn is applied to the element v to determine if it should be included in the result.
		if fn(v) {
			// If the element satisfies the condition, it is added to the result slice.
			// The append function dynamically grows the result slice and appends the element v to it.
			result = append(result, v)
		}
	}

	// Return the resulting slice containing only the elements that satisfy the predicate.
	return result
}

// Unique removes duplicate elements from a slice of any comparable type.
// It iterates over each element in the input slice and keeps track of the elements that have already been encountered.
// If an element has not been encountered before, it is added to the result slice.
// The result is a new slice containing only the unique elements, preserving their original order.
// This function is generic and works with any comparable type, including integers, strings, structs, and more.
func Unique[T comparable](elements []T) []T {
	// Declare an empty slice to hold the unique elements.
	// The result slice will store the final list of elements with duplicates removed.
	var result []T
	// Create a map to track the elements that have already been encountered.
	// The map's key is the element itself, and the value is a boolean indicating if the element has been seen.
	// The map's capacity is set to the length of the input slice for optimization.
	seen := make(map[T]bool, len(elements))

	// Iterate over each element in the input slice.
	// The variable 'elem' holds the current element being processed from the input slice.
	for _, elem := range elements {
		// Check if the current element has already been seen.
		// The 'ok' variable checks if the element exists in the 'seen' map.
		if ok := seen[elem]; ok {
			// If the element has been seen, skip it by continuing to the next iteration.
			// This prevents duplicates from being added to the result.
			continue
		}

		// If the element has not been seen before, mark it as seen.
		// This is done by setting the value for the element in the 'seen' map to 'true'.
		seen[elem] = true

		// Append the unique element to the result slice.
		// This action ensures that only unique elements are included in the final result.
		result = append(result, elem)
	}

	// Return the result slice containing only unique elements from the input slice.
	// The order of the elements is preserved.
	return result
}
