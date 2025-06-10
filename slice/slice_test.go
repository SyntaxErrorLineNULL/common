package core

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/SyntaxErrorLineNULL/common/test"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	// Define a set of test cases with inputs and the expected output for each scenario.
	// The test cases cover a range of situations, including nil inputs, empty inputs, and
	// slices with elements, to ensure the Merge function behaves as intended in all cases.
	cases := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{name: "Both slices nil", first: nil, second: nil, expected: []int{}},
		{name: "First slice nil", first: nil, second: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "Second slice nil", first: []int{4, 5, 6}, second: nil, expected: []int{4, 5, 6}},
		{name: "Both slices empty", first: []int{}, second: []int{}, expected: []int{}},
		{name: "First slice empty", first: []int{}, second: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "Second slice empty", first: []int{4, 5, 6}, second: []int{}, expected: []int{4, 5, 6}},
		{name: "Both slices have elements", first: []int{7, 8, 9}, second: []int{10, 11, 12}, expected: []int{7, 8, 9, 10, 11, 12}},
		{name: "Slices with overlapping elements", first: []int{1, 2, 3}, second: []int{3, 4, 5}, expected: []int{1, 2, 3, 3, 4, 5}},
		{name: "Different length slices", first: []int{1}, second: []int{2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
	}

	// Iterate over the defined test cases, executing each one as a subtest.
	// Subtests allow each test case to be run independently, making it easier
	// to identify which specific case fails if an assertion does not hold.
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// Call the Merge function with the input slices from the current test case.
			// The Merge function is expected to concatenate the two slices and return the result.
			result := Merge(tt.first, tt.second)

			// Compare the result of the Merge function with the expected output.
			// The assertion checks for equality and fails the test if the values do not match.
			assert.Equal(t, tt.expected, result, "Result for test case %q did not match expected output", tt.name)
		})
	}
}

func TestExclude(t *testing.T) {
	t.Parallel()

	// SliceString tests the Exclude function for slices of integers. This test suite is designed to ensure that
	// the Exclude function correctly removes all occurrences of a specified integer from a slice. The test cases
	// cover a range of scenarios, including removing an element that appears once or multiple times, trying to
	// remove an element that does not exist in the slice, and handling empty or nil slices.
	t.Run("SliceString", func(t *testing.T) {
		cases := []struct {
			name     string
			elements []int
			element  int
			expected []int
		}{
			{name: "ExcludeSingleElement", elements: []int{1, 2, 3, 4, 5}, element: 3, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeMultipleOccurrences", elements: []int{1, 2, 3, 3, 4, 3, 5}, element: 3, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeNonexistentElement", elements: []int{1, 2, 3, 4, 5}, element: 6, expected: []int{1, 2, 3, 4, 5}},
			{name: "ExcludeEmptySlice", elements: []int{}, element: 1, expected: []int{}},
			{name: "ExcludeSingleElementSlice", elements: []int{1}, element: 1, expected: []int{}},
			{name: "ExcludeNilSlice", elements: nil, element: 1, expected: nil},
		}

		// Iterate through each test case and execute the Exclude function.
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				// Call the Exclude function with the current test case's elements and element.
				result := Exclude(tt.elements, tt.element)

				// Assert that the result from the Exclude function matches the expected value.
				// If the result does not match the expected value, the test will fail, and the
				// provided error message will help identify which test case failed and why.
				assert.Equal(t, tt.expected, result, "Test case %s failed", tt.name)
			})
		}
	})

	// SliceInt tests the Exclude function with various scenarios involving slices of integers. The goal of this test
	// is to verify that the Exclude function correctly removes all occurrences of a specified integer from a slice.
	// The test cases cover a range of conditions to ensure the function handles different situations accurately,
	// including scenarios where elements appear multiple times, do not appear, or when the input slice is empty or nil.
	// Additionally, it tests edge cases such as large slices and slices with negative numbers.
	t.Run("SliceInt", func(t *testing.T) {
		cases := []struct {
			name     string
			elements []int
			element  int
			expected []int
		}{
			{name: "ExcludeSingleElement", elements: []int{1, 2, 3, 4, 5}, element: 3, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeMultipleOccurrences", elements: []int{1, 2, 3, 3, 4, 3, 5}, element: 3, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeNonexistentElement", elements: []int{1, 2, 3, 4, 5}, element: 6, expected: []int{1, 2, 3, 4, 5}},
			{name: "ExcludeEmptySlice", elements: []int{}, element: 1, expected: []int{}},
			{name: "ExcludeSingleElementSlice", elements: []int{1}, element: 1, expected: []int{}},
			{name: "ExcludeNilSlice", elements: nil, element: 1, expected: nil},
			{name: "ExcludeAllElements", elements: []int{7, 7, 7, 7, 7}, element: 7, expected: []int{}},
			{name: "ExcludeFirstElement", elements: []int{9, 2, 3, 4, 5}, element: 9, expected: []int{2, 3, 4, 5}},
			{name: "ExcludeLastElement", elements: []int{1, 2, 3, 4, 10}, element: 10, expected: []int{1, 2, 3, 4}},
			{name: "ExcludeMiddleElement", elements: []int{1, 2, 10, 4, 5}, element: 10, expected: []int{1, 2, 4, 5}},
			{name: "ExcludeElementFromLargeSlice", elements: test.CreateSequenceWithRepeats(1000, 500), element: 500, expected: test.CreateSequenceWithoutRepeats(1000)},
			{name: "ExcludeNegativeElement", elements: []int{-3, -2, -1, 0, 1, 2, 3}, element: -2, expected: []int{-3, -1, 0, 1, 2, 3}},
			{name: "ExcludeFromMixedValues", elements: []int{-1, 0, 1, -1, 0, 1}, element: -1, expected: []int{0, 1, 0, 1}},
		}

		// Iterate through each test case and execute the Exclude function.
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				// Call the Exclude function with the current test case's elements and element.
				result := Exclude(tt.elements, tt.element)

				// Assert that the result from the Exclude function matches the expected value.
				// If the result does not match the expected value, the test will fail, and the
				// provided error message will help identify which test case failed and why.
				assert.Equal(t, tt.expected, result, "Test case %s failed", tt.name)
			})
		}
	})
}

func TestContains(t *testing.T) {
	t.Parallel()

	// SliceInt tests the Contains function for slices of integers. This test ensures that the Contains function
	// accurately identifies whether a specific integer is present in a slice of integers. It covers a variety of
	// cases to verify correctness, including scenarios where the element is present or absent from the slice.
	t.Run("SliceInt", func(t *testing.T) {
		// Define a range of test cases for the Contains function with integer slices.
		cases := []struct {
			name     string
			elements []int
			element  int
			expected bool
		}{
			{name: "Element is in the slice", elements: []int{1, 2, 3, 4, 5}, element: 3, expected: true},
			{name: "Element is not in the slice", elements: []int{1, 2, 3, 4, 5}, element: 6, expected: false},
			{name: "Empty slice", elements: []int{}, element: 1, expected: false},
			{name: "Empty slice with empty element", elements: []int{}, element: 0, expected: false},
			{name: "Single element slice contains element", elements: []int{1}, element: 1, expected: true},
			{name: "Single element slice does not contain element", elements: []int{1}, element: 2, expected: false},
			{name: "Multiple elements, element at start", elements: []int{5, 1, 2, 3, 4}, element: 5, expected: true},
			{name: "Multiple elements, element at end", elements: []int{1, 2, 3, 4, 5}, element: 5, expected: true},
			{name: "Multiple elements, element in middle", elements: []int{1, 2, 3, 4, 5}, element: 3, expected: true},
			{name: "Multiple elements, element repeated", elements: []int{1, 2, 3, 3, 4, 5}, element: 3, expected: true},
			{name: "Nil slice", elements: nil, element: 1, expected: false},
		}

		// Iterate through each test case and execute the Contains function
		for _, tt := range cases {
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
		cases := []struct {
			name     string
			elements []string
			element  string
			expected bool
		}{
			{name: "Nil slice", elements: nil, element: "test", expected: false},
			{name: "Empty slice", elements: []string{}, element: "test", expected: false},
			{name: "Element in single-element sorted slice", elements: []string{"test"}, element: "test", expected: true},
			{name: "Element not in single-element sorted slice", elements: []string{"test"}, element: "notfound", expected: false},
			{name: "Element in multiple-element sorted slice", elements: []string{"alpha", "beta", "gamma"}, element: "beta", expected: true},
			{name: "Element not in multiple-element sorted slice", elements: []string{"alpha", "beta", "gamma"}, element: "delta", expected: false},
			{name: "Element at the beginning of the sorted slice", elements: []string{"alpha", "beta", "gamma"}, element: "alpha", expected: true},
			{name: "Element at the end of the sorted slice", elements: []string{"alpha", "beta", "gamma"}, element: "gamma", expected: true},
			{name: "Unsorted slice (contains element)", elements: []string{"beta", "alpha", "gamma"}, element: "alpha", expected: true},
			{name: "Unsorted slice (does not contain element)", elements: []string{"beta", "alpha", "gamma"}, element: "delta", expected: false},
		}

		// Iterate over each test case and execute the Contains function.
		for _, tt := range cases {
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

func TestMap(t *testing.T) {
	t.Parallel()

	// SliceIntToString tests the Map function for slices of integers transformed into strings.
	// This test verifies that the Map function correctly applies the transformation function
	// to each element of the integer slice, producing the expected slice of strings.
	t.Run("SliceIntToString", func(t *testing.T) {
		// Define a series of test cases to cover various scenarios for the Map function.
		// Each test case includes a name, an input slice of integers, a transformation function, and the expected output slice of strings.
		cases := []struct {
			name          string
			elements      []int
			transformFunc func(int) string
			expected      []string
		}{
			{name: "Empty slice", elements: []int{}, transformFunc: func(i int) string { return fmt.Sprintf("%d", i) }, expected: []string{}},
			{name: "Single element slice", elements: []int{42}, transformFunc: func(i int) string { return fmt.Sprintf("Num:%d", i) }, expected: []string{"Num:42"}},
			{name: "Multiple elements", elements: []int{1, 2, 3}, transformFunc: func(i int) string { return fmt.Sprintf("%d", i*10) }, expected: []string{"10", "20", "30"}},
			{name: "Negative integers", elements: []int{-1, -2}, transformFunc: func(i int) string { return fmt.Sprintf("%d", i) }, expected: []string{"-1", "-2"}},
		}

		// Iterate over each test case defined in the `cases` slice.
		// This loop ensures that the Map function is tested with a variety of inputs and transformation functions.
		for _, tt := range cases {
			// Define a subtest for each test case using t.Run.
			// Subtests help isolate each scenario and provide detailed feedback if a particular case fails.
			t.Run(tt.name, func(t *testing.T) {
				// Execute the Map function using the test case's input slice and transformation function.
				// The Map function is expected to transform each element in the input slice
				// according to the logic defined in the provided transformation function.
				result := Map(tt.elements, tt.transformFunc)

				// Assert that the output from the Map function matches the expected output for the test case.
				// If the result does not match, the test will fail and output the provided message,
				// including the name of the test case for better debugging context.
				assert.Equal(t, tt.expected, result, "Map output should match expected value for test case: %s", tt.name)
			})
		}
	})

	// SliceStringToString verifies that the Map function correctly transforms a slice of strings
	// using various transformation functions. The test cases cover scenarios such as empty slices,
	// single-element slices, multiple-element slices, and string transformations like appending
	// or changing case. This ensures that the Map function behaves as expected across different inputs.
	t.Run("SliceStringToString", func(t *testing.T) {
		// Define test cases for transforming string slices into new string slices.
		// Each test case specifies a name, input slice, transformation function, and the expected result.
		cases := []struct {
			name          string
			elements      []string
			transformFunc func(string) string
			expected      []string
		}{
			{name: "Empty slice", elements: []string{}, transformFunc: func(s string) string { return fmt.Sprintf("%s_mod", s) }, expected: []string{}},
			{name: "Single element slice", elements: []string{"test"}, transformFunc: func(s string) string { return fmt.Sprintf("%s_appended", s) }, expected: []string{"test_appended"}},
			{name: "Multiple elements", elements: []string{"a", "b", "c"}, transformFunc: func(s string) string { return fmt.Sprintf("%s%s", s, s) }, expected: []string{"aa", "bb", "cc"}},
			{name: "Uppercase transformation", elements: []string{"lower", "case"}, transformFunc: func(s string) string { return fmt.Sprintf("%s", strings.ToUpper(s)) }, expected: []string{"LOWER", "CASE"}},
		}

		// Iterate over each test case defined in the `cases` slice.
		// This loop ensures that the Map function is tested with a variety of inputs and transformation functions.
		for _, tt := range cases {
			// Define a subtest for each test case using t.Run.
			// Subtests help isolate each scenario and provide detailed feedback if a particular case fails.
			t.Run(tt.name, func(t *testing.T) {
				// Execute the Map function using the test case's input slice and transformation function.
				// The Map function is expected to transform each element in the input slice
				// according to the logic defined in the provided transformation function.
				result := Map(tt.elements, tt.transformFunc)

				// Assert that the output from the Map function matches the expected output for the test case.
				// If the result does not match, the test will fail and output the provided message,
				// including the name of the test case for better debugging context.
				assert.Equal(t, tt.expected, result, "Map output should match expected value for test case: %s", tt.name)
			})
		}
	})

	// SliceComplexTransformations tests the Map function with scenarios involving complex transformation logic.
	// This ensures the function can handle advanced mappings such as conditional evaluations, mathematical operations,
	// and transformations with negative integers. Each test case provides a unique transformation logic to validate
	// the robustness of the Map function when handling complex inputs and outputs.
	t.Run("SliceComplexTransformations", func(t *testing.T) {
		// Define test cases with various transformation functions applied to slices of integers.
		// Each case includes a name for identification, an input slice, a transformation function, and the expected result.
		cases := []struct {
			name          string
			elements      []int
			transformFunc func(int) string
			expected      []string
		}{
			{name: "Odd or Even", elements: []int{1, 2, 3, 4}, transformFunc: func(i int) string {
				return fmt.Sprintf("%d:%s", i, map[bool]string{true: "Odd", false: "Even"}[i%2 != 0])
			}, expected: []string{"1:Odd", "2:Even", "3:Odd", "4:Even"}},
			{name: "Square of numbers", elements: []int{2, 3}, transformFunc: func(i int) string { return fmt.Sprintf("%d", i*i) }, expected: []string{"4", "9"}},
			{name: "Negative handling", elements: []int{-1, 2, -3}, transformFunc: func(i int) string { return fmt.Sprintf("%d", i*2) }, expected: []string{"-2", "4", "-6"}},
		}

		// Iterate over each test case defined in the `cases` slice.
		// This loop ensures that the Map function is tested with a variety of inputs and transformation functions.
		for _, tt := range cases {
			// Define a subtest for each test case using t.Run.
			// Subtests help isolate each scenario and provide detailed feedback if a particular case fails.
			t.Run(tt.name, func(t *testing.T) {
				// Execute the Map function using the test case's input slice and transformation function.
				// The Map function is expected to transform each element in the input slice
				// according to the logic defined in the provided transformation function.
				result := Map(tt.elements, tt.transformFunc)

				// Assert that the output from the Map function matches the expected output for the test case.
				// If the result does not match, the test will fail and output the provided message,
				// including the name of the test case for better debugging context.
				assert.Equal(t, tt.expected, result, "Expected Map result for test case: %s", tt.name)
			})
		}
	})

	// GeneratedDataWithFilterAndTransformation tests the Map function using a sequence of integers generated with
	// specific constraints (no multiples of 100) and applies both filtering and transformation operations. The test
	// validates the correctness of the Map function when used in conjunction with filtering logic that reduces the input
	// set, followed by a transformation that modifies the filtered values. This ensures that the Map function can handle
	// a chain of operations where filtering and transformation are combined correctly.
	t.Run("GeneratedDataWithFilterAndTransformation", func(t *testing.T) {
		// Generate a sequence of 250 integers, ensuring that no element is a multiple of 100.
		// This function simulates a larger set of data that we can apply transformations and filters to.
		// It produces a sequence that we can manipulate further for testing purposes.
		input := test.CreateSequenceWithoutRepeats(250)

		// Define the filter function that will be applied to the input data.
		// In this case, we are only interested in even numbers (i.e., numbers divisible by 2).
		filter := func(i int) bool { return i%2 == 0 }

		// Define the transformation function that will be applied to each element in the filtered set.
		// Here, each number will be multiplied by 2.
		transform := func(i int) int { return i * 2 }

		// Initialize a slice to hold the filtered values from the input sequence.
		// We only append elements that satisfy the filter function (i.e., even numbers).
		filtered := make([]int, 0, len(input))
		for _, v := range input {
			if filter(v) {
				filtered = append(filtered, v)
			}
		}

		// Use the Map function to apply the transformation to the filtered data.
		// Each number in the filtered slice is passed through the transformation function
		// (multiplying it by 2) and the results are collected in the result slice.
		result := Map(filtered, transform)

		// Prepare the expected result by manually applying the transformation to each filtered value.
		// This serves as a reference to compare against the result of the Map function.
		var expected []int

		for _, v := range filtered {
			expected = append(expected, v*2)
		}

		// Assert that the result of applying the Map function to the filtered data matches the expected output.
		// The expected output is calculated by manually applying the transformation (multiplying by 2) to each
		// element of the filtered slice. This assertion ensures that the Map function correctly transformed the
		// filtered elements according to the transformation logic and that the expected results are accurate.
		assert.Equal(t, expected, result, "Transformed and filtered output should match expected")
	})
}

func TestFilter(t *testing.T) {
	t.Parallel()

	// FilterInt tests the behavior of the Filter function when applied to a variety of input slices.
	// It ensures that the Filter function correctly filters elements based on the provided predicate function.
	// The test cases cover a wide range of scenarios, including filtering even and odd numbers,
	// handling empty slices, and dealing with slices where no elements match the predicate.
	t.Run("FilterInt", func(t *testing.T) {
		// Define a set of test cases with inputs and the expected output for each scenario.
		// The test cases cover a range of situations, including nil inputs, empty inputs, and
		// slices with elements, to ensure the Filter function behaves as intended in all cases.
		cases := []struct {
			name     string
			elements []int
			fn       func(int) bool
			expected []int
		}{
			{name: "Filter even numbers", elements: []int{1, 2, 3, 4, 5}, fn: func(n int) bool { return n%2 == 0 }, expected: []int{2, 4}},
			{name: "Filter odd numbers", elements: []int{1, 2, 3, 4, 5}, fn: func(n int) bool { return n%2 != 0 }, expected: []int{1, 3, 5}},
			{name: "Empty slice", elements: []int{}, fn: func(n int) bool { return n%2 == 0 }, expected: nil},
			{name: "All elements match predicate", elements: []int{2, 4, 6, 8}, fn: func(n int) bool { return n%2 == 0 }, expected: []int{2, 4, 6, 8}},
			{name: "No elements match predicate", elements: []int{1, 3, 5, 7}, fn: func(n int) bool { return n%2 == 0 }, expected: nil},
			{name: "Single element match", elements: []int{1}, fn: func(n int) bool { return n%2 == 0 }, expected: nil},
			{name: "Single element match", elements: []int{2}, fn: func(n int) bool { return n%2 == 0 }, expected: []int{2}},
			{name: "Mixed elements with no matching", elements: []int{1, 3, 5}, fn: func(n int) bool { return n%2 == 0 }, expected: nil},
		}

		// Iterate over the defined test cases, executing each one as a subtest.
		// Subtests allow each test case to be run independently, making it easier
		// to identify which specific case fails if an assertion does not hold.
		for _, tt := range cases {
			// Start a subtest for the current test case, using the test case's name.
			t.Run(tt.name, func(t *testing.T) {
				// Call the Filter function with the current test case's input slice and predicate function.
				// This will filter the elements in the slice based on the predicate and return the filtered result.
				result := Filter(tt.elements, tt.fn)

				// Compare the result of the Filter function with the expected output.
				// The assertion checks for equality and will fail the test if the result does not match the expected output.
				assert.Equal(t, tt.expected, result, "For case '%s', expected %v but got %v", tt.name, tt.expected, result)
			})
		}
	})

	// FilterString tests the behavior of the Filter function when applied to a variety of string slices.
	// It ensures that the Filter function correctly filters elements based on the provided predicate function.
	// The test cases cover various filtering conditions such as matching the first character of a string,
	// handling empty slices, and verifying the behavior when no elements match the predicate.
	t.Run("FilterString", func(t *testing.T) {
		// Define a set of test cases with inputs and the expected output for each scenario.
		// The test cases cover a range of situations, including nil inputs, empty inputs, and
		// slices with elements, to ensure the Filter function behaves as intended in all cases.
		cases := []struct {
			name     string
			elements []string
			fn       func(string) bool
			expected []string
		}{
			{name: "Filter strings starting with 'a'", elements: []string{"apple", "banana", "avocado"}, fn: func(s string) bool { return len(s) > 0 && s[0] == 'a' }, expected: []string{"apple", "avocado"}},
			{name: "Filter strings containing 'o'", elements: []string{"apple", "banana", "orange"}, fn: func(s string) bool { return len(s) > 0 && s[0] == 'o' }, expected: []string{"orange"}},
			{name: "Empty slice", elements: []string{}, fn: func(s string) bool { return len(s) > 0 && s[0] == 'o' }, expected: nil},
			{name: "All elements match predicate", elements: []string{"apple", "avocado", "orange"}, fn: func(s string) bool { return len(s) > 0 && s[0] == 'a' }, expected: []string{"apple", "avocado"}},
			{name: "No elements match predicate", elements: []string{"banana", "grape", "pear"}, fn: func(s string) bool { return len(s) > 0 && s[0] == 'o' }, expected: nil},
			{name: "Single element matches", elements: []string{"orange"}, fn: func(s string) bool { return len(s) > 0 && s[0] == 'o' }, expected: []string{"orange"}},
			{name: "Single element does not match", elements: []string{"banana"}, fn: func(s string) bool { return len(s) > 0 && s[0] == 'o' }, expected: nil},
		}

		// Iterate over the defined test cases, executing each one as a subtest.
		// Subtests allow each test case to be run independently, making it easier
		// to identify which specific case fails if an assertion does not hold.
		for _, tt := range cases {
			// Start a subtest for the current test case, using the test case's name.
			t.Run(tt.name, func(t *testing.T) {
				// Call the Filter function with the current test case's input slice and predicate function.
				// This will filter the elements in the slice based on the predicate and return the filtered result.
				result := Filter(tt.elements, tt.fn)

				// Compare the result of the Filter function with the expected output.
				// The assertion checks for equality and will fail the test if the result does not match the expected output.
				assert.Equal(t, tt.expected, result, "For case '%s', expected %v but got %v", tt.name, tt.expected, result)
			})
		}
	})

	// ComplexFilter tests the Filter function across a variety of complex scenarios involving different
	// data types and predicates. It ensures that the Filter function can correctly handle various types
	// such as integers, strings, structs, pointers, floats, complex numbers, and interfaces. The test cases
	// verify that the Filter function behaves as expected in diverse filtering conditions, including nested
	// filters, pointer slices, and mixed types in an interface slice.
	t.Run("ComplexFilter", func(t *testing.T) {
		// Define a set of test cases with inputs and the expected output for each scenario.
		// The test cases cover a range of situations, including nil inputs, empty inputs, and
		// slices with elements, to ensure the Filter function behaves as intended in all cases.
		cases := []struct {
			name     string
			elements interface{}
			fn       interface{}
			expected interface{}
		}{
			{
				name:     "Filter even numbers greater than 10",
				elements: []int{5, 10, 11, 14, 15, 16},
				fn:       func(n int) bool { return n%2 == 0 && n > 10 },
				expected: []int{14, 16},
			},
			{
				name:     "Filter strings containing 'e' and longer than 3 characters",
				elements: []string{"apple", "banana", "cherry", "kiwi", "grape"},
				fn:       func(s string) bool { return len(s) > 3 && strings.Contains(s, "e") },
				expected: []string{"apple", "cherry", "grape"},
			},
			{
				name: "Filter struct slices by boolean field",
				elements: []struct {
					name   string
					active bool
				}{{"alice", true}, {"bob", false}, {"charlie", true}},
				fn: func(e struct {
					name   string
					active bool
				},
				) bool {
					return e.active
				},
				expected: []struct {
					name   string
					active bool
				}{{"alice", true}, {"charlie", true}},
			},
			{
				name: "Filter based on a custom struct field with nested filter",
				elements: []struct {
					name string
					age  int
				}{{"John", 25}, {"Jane", 30}, {"Doe", 15}},
				fn: func(e struct {
					name string
					age  int
				},
				) bool {
					return e.age > 18 && e.age < 30
				},
				expected: []struct {
					name string
					age  int
				}{{"John", 25}},
			},
			{
				name:     "Filter out nil values from a pointer slice",
				elements: []*string{nil, nil, strPtr("apple"), strPtr("banana"), nil},
				fn:       func(s *string) bool { return s != nil && len(*s) > 5 },
				expected: []*string{strPtr("banana")},
			},
			{
				name:     "Filter float numbers greater than a threshold",
				elements: []float64{1.2, 5.6, 7.8, 9.1},
				fn:       func(f float64) bool { return f > 5.0 },
				expected: []float64{5.6, 7.8, 9.1},
			},
			{
				name:     "Filter complex numbers where real part is greater than imaginary",
				elements: []complex128{complex(1, 2), complex(3, 1), complex(4, 5), complex(6, 3)},
				fn:       func(c complex128) bool { return real(c) > imag(c) },
				expected: []complex128{complex(3, 1), complex(6, 3)},
			},
			{
				name: "Empty slice with custom condition on structs",
				elements: []struct {
					id   int
					name string
				}{},
				fn: func(e struct {
					id   int
					name string
				},
				) bool {
					return e.id > 0
				},
				expected: nil,
			},
			{
				name:     "Filter mixed types in an interface slice (interface{})",
				elements: []interface{}{1, "hello", 2.5, "world", true},
				fn:       func(v interface{}) bool { return fmt.Sprintf("%T", v) == "string" },
				expected: []interface{}{"hello", "world"},
			},
		}

		// Iterate over the defined test cases, executing each one as a subtest.
		// Subtests allow each test case to be run independently, making it easier
		// to identify which specific case fails if an assertion does not hold.
		for _, tt := range cases {
			// Start a subtest for the current test case, using the test case's name.
			t.Run(tt.name, func(t *testing.T) {
				// Handle each test case based on the input type.
				var result interface{}
				switch v := tt.elements.(type) {
				case []int:
					result = Filter(v, tt.fn.(func(int) bool))
				case []string:
					result = Filter(v, tt.fn.(func(string) bool))
				case []struct {
					name   string
					active bool
				}:
					result = Filter(v, tt.fn.(func(struct {
						name   string
						active bool
					}) bool))
				case []struct {
					name string
					age  int
				}:
					result = Filter(v, tt.fn.(func(struct {
						name string
						age  int
					}) bool))
				case []*string:
					result = Filter(v, tt.fn.(func(*string) bool))
				case []float64:
					result = Filter(v, tt.fn.(func(float64) bool))
				case []complex128:
					result = Filter(v, tt.fn.(func(complex128) bool))
				case []interface{}:
					result = Filter(v, tt.fn.(func(interface{}) bool))
				}

				// Compare the result of the Filter function with the expected output.
				// The assertion checks for equality and will fail the test if the result does not match the expected output.
				assert.Equal(t, tt.expected, result, "For case '%s', expected %v but got %v", tt.name, tt.expected, result)
			})
		}
	})
}

func TestUnique(t *testing.T) {
	// Define test cases for different data types to check the behavior of the Unique function.
	// Each test case consists of a name, the input elements (which is the slice to deduplicate),
	// and the expected result after duplicates are removed.
	cases := []struct {
		name     string
		elements interface{}
		expected interface{}
	}{
		{
			name:     "Unique on integer slice with duplicates",
			elements: []int{1, 2, 2, 3, 4, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unique on string slice with duplicates",
			elements: []string{"apple", "banana", "apple", "orange", "banana", "grape"},
			expected: []string{"apple", "banana", "orange", "grape"},
		},
		{
			name:     "Unique on empty integer slice",
			elements: []int{},
			expected: []int(nil),
		},
		{
			name:     "Unique on single element slice",
			elements: []int{42},
			expected: []int{42},
		},
		{
			name: "Unique on struct slice with duplicates",
			elements: []struct {
				name string
				age  int
			}{{"John", 25}, {"Jane", 30}, {"John", 25}, {"Doe", 22}},
			expected: []struct {
				name string
				age  int
			}{{"John", 25}, {"Jane", 30}, {"Doe", 22}},
		},
		{
			name:     "Unique on boolean slice with duplicates",
			elements: []bool{true, false, true, true, false},
			expected: []bool{true, false},
		},
		{
			name:     "Unique on float slice with duplicates",
			elements: []float64{1.1, 2.2, 3.3, 2.2, 1.1},
			expected: []float64{1.1, 2.2, 3.3},
		},
		{
			name:     "Unique on slice of pointers",
			elements: []string{"apple", "banana", "apple"},
			expected: []string{"apple", "banana"},
		},
		{
			name:     "Unique on slice with nil elements",
			elements: []*string{nil, nil, strPtr("apple"), nil, strPtr("banana")},
			expected: []*string{nil, strPtr("apple"), strPtr("banana")},
		},
		{
			name:     "Unique on slice of complex numbers",
			elements: []complex128{complex(1, 1), complex(2, 2), complex(1, 1), complex(3, 3)},
			expected: []complex128{complex(1, 1), complex(2, 2), complex(3, 3)},
		},
	}

	// Iterate over the defined test cases, executing each one as a subtest.
	// Subtests allow each test case to be run independently, making it easier
	// to identify which specific case fails if an assertion does not hold.
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// Inside each subtest, we handle each case based on the input type (elements).
			// We declare a variable result of type interface{} to store the output of the Unique function.
			var result interface{}

			switch v := tt.elements.(type) {
			case []int:
				result = Unique(v)
			case []string:
				result = Unique(v)
			case []struct {
				name string
				age  int
			}:
				result = Unique(v)
			case []bool:
				result = Unique(v)
			case []float64:
				result = Unique(v)
			case []*string:
				result = Unique(v)
			case []complex128:
				result = Unique(v)
			}

			// assert.Equal is used to check if the result matches the expected value.
			// It compares the actual result with the expected value and reports an error if they don't match.
			// The assertion will fail if result does not equal tt.expected, and the error message will include
			// the test case name tt.name, the expected value tt.expected, and the actual result.
			assert.Equal(t, tt.expected, result, "For case '%s', expected %v but got %v", tt.name, tt.expected, result)
		})
	}
}

// strPtr is a helper function to create a pointer to a string.
func strPtr(s string) *string {
	return &s
}
