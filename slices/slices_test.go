package slices

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMerge is a table-driven test function for testing the Merge function.
// It iterates over a series of test cases, each with different input slices (first, second),
// and checks whether the Merge function produces the expected result (expected).
func TestMerge(t *testing.T) {
	// Define a slice of test cases.
	cases := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "Both slices nil",
			first:    nil,
			second:   nil,
			expected: []int{},
		},
		{
			name:     "First slice nil",
			first:    nil,
			second:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Second slice nil",
			first:    []int{4, 5, 6},
			second:   nil,
			expected: []int{4, 5, 6},
		},
		{
			name:     "Both slices empty",
			first:    []int{},
			second:   []int{},
			expected: []int{},
		},
		{
			name:     "First slice empty",
			first:    []int{},
			second:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Second slice empty",
			first:    []int{4, 5, 6},
			second:   []int{},
			expected: []int{4, 5, 6},
		},
		{
			name:     "Both slices have elements",
			first:    []int{7, 8, 9},
			second:   []int{10, 11, 12},
			expected: []int{7, 8, 9, 10, 11, 12},
		},
		{
			name:     "Slices with overlapping elements",
			first:    []int{1, 2, 3},
			second:   []int{3, 4, 5},
			expected: []int{1, 2, 3, 3, 4, 5},
		},
		{
			name:     "Different length slices",
			first:    []int{1},
			second:   []int{2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	// Iterate over each test case.
	for _, tt := range cases {
		// Run each test case as a subtest.
		t.Run(tt.name, func(t *testing.T) {
			// Call the Merge function with the current test case's input slices.
			result := Merge(tt.first, tt.second)
			// Use assert.Equal to check if the output matches the expected result.
			// If not, the test will fail with a message showing the difference.
			assert.Equal(t, tt.expected, result, "they should be equal")
		})
	}
}

func TestContains(t *testing.T) {
	t.Parallel()

	// SliceInt tests the Contains function for slices of integers. This test ensures that the Contains function
	// accurately identifies whether a specific integer is present in a slice of integers. It covers a variety of
	// cases to verify correctness, including scenarios where the element is present or absent from the slice.
	t.Run("SliceInt", func(t *testing.T) {
		// Define a range of test cases for the Contains function with integer slices.
		tests := []struct {
			name     string
			elements []int
			element  int
			expected bool
		}{
			{
				name:     "Element is in the slice",
				elements: []int{1, 2, 3, 4, 5},
				element:  3,
				expected: true,
			},
			{
				name:     "Element is not in the slice",
				elements: []int{1, 2, 3, 4, 5},
				element:  6,
				expected: false,
			},
			{
				name:     "Empty slice",
				elements: []int{},
				element:  1,
				expected: false,
			},
			{
				name:     "Empty slice with empty element",
				elements: []int{},
				element:  0,
				expected: false,
			},
			{
				name:     "Single element slice contains element",
				elements: []int{1},
				element:  1,
				expected: true,
			},
			{
				name:     "Single element slice does not contain element",
				elements: []int{1},
				element:  2,
				expected: false,
			},
			{
				name:     "Multiple elements, element at start",
				elements: []int{5, 1, 2, 3, 4},
				element:  5,
				expected: true,
			},
			{
				name:     "Multiple elements, element at end",
				elements: []int{1, 2, 3, 4, 5},
				element:  5,
				expected: true,
			},
			{
				name:     "Multiple elements, element in middle",
				elements: []int{1, 2, 3, 4, 5},
				element:  3,
				expected: true,
			},
			{
				name:     "Multiple elements, element repeated",
				elements: []int{1, 2, 3, 3, 4, 5},
				element:  3,
				expected: true,
			},
			{
				name:     "Nil slice",
				elements: nil,
				element:  1,
				expected: false,
			},
		}

		// Iterate through each test case and execute the Contains function
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Call the Contains function with the current test case's elements and element.
				// The Contains function will return whether the element is present in the slice.
				result := Contains(tt.elements, tt.element)

				// Assert that the result from the Contains function matches the expected value.
				// If the result does not match the expected value, the test will fail, and the
				// provided error message will help identify which test case failed and why.
				assert.Equal(t, tt.expected, result, "result should match the expected value for test case: %s", tt.name)
			})
		}
	})

	// SliceString tests the Contains function for slices of strings.
	// It verifies that the function correctly identifies the presence or absence of an element
	// in various scenarios involving string slices, including sorted and unsorted slices.
	t.Run("SliceString", func(t *testing.T) {
		// Define test cases with various scenarios for slices of strings.
		tests := []struct {
			name     string
			elements []string
			element  string
			expected bool
		}{
			{
				name:     "Nil slice",
				elements: nil,
				element:  "test",
				expected: false,
			},
			{
				name:     "Empty slice",
				elements: []string{},
				element:  "test",
				expected: false,
			},
			{
				name:     "Element in single-element sorted slice",
				elements: []string{"test"},
				element:  "test",
				expected: true,
			},
			{
				name:     "Element not in single-element sorted slice",
				elements: []string{"test"},
				element:  "notfound",
				expected: false,
			},
			{
				name:     "Element in multiple-element sorted slice",
				elements: []string{"alpha", "beta", "gamma"},
				element:  "beta",
				expected: true,
			},
			{
				name:     "Element not in multiple-element sorted slice",
				elements: []string{"alpha", "beta", "gamma"},
				element:  "delta",
				expected: false,
			},
			{
				name:     "Element at the beginning of the sorted slice",
				elements: []string{"alpha", "beta", "gamma"},
				element:  "alpha",
				expected: true,
			},
			{
				name:     "Element at the end of the sorted slice",
				elements: []string{"alpha", "beta", "gamma"},
				element:  "gamma",
				expected: true,
			},
			{
				name:     "Unsorted slice (contains element)",
				elements: []string{"beta", "alpha", "gamma"},
				element:  "alpha",
				expected: true,
			},
			{
				name:     "Unsorted slice (does not contain element)",
				elements: []string{"beta", "alpha", "gamma"},
				element:  "delta",
				expected: false,
			},
		}

		// Iterate over each test case and execute the Contains function.
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Sort the slice of strings to prepare it for testing the Contains function.
				// Sorting ensures that the Contains function operates correctly on ordered data.
				// This step is essential if the slice is not nil, as it standardizes the input for the test.
				if tt.elements != nil {
					// Sort the slice of strings in ascending order.
					// Sorting is done to match the expected behavior of the Contains function when the slice is ordered.
					sort.Strings(tt.elements)
				}

				// Call the Contains function with the current test case’s slice and element.
				result := Contains(tt.elements, tt.element)

				// Assert that the result matches the expected value.
				// If the result does not match, the test will fail, and the message will include
				// the test case details for easy identification of the failure.
				assert.Equal(t, tt.expected, result, "Expected Contains(%v, %v) to be %v but result %v", tt.elements, tt.element, tt.expected, result)
			})
		}
	})
}
