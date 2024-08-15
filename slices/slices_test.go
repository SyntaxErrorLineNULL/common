package slices

import (
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
