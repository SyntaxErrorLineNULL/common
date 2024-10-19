package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSplitStringBySeparator verifies the behavior of the SplitStringBySeparator function.
// The test ensures that the function correctly splits an input string based on the given separator.
// It covers various scenarios, such as when the separator is present in the middle, at the beginning,
// or at the end of the string, as well as cases where the separator is missing or empty.
// The test checks the correctness of the split by asserting that the returned before and after parts
// match the expected values and verifies whether the function correctly identifies if the separator was found.
func TestSplitStringBySeparator(t *testing.T) {
	// Define a slice of test cases to cover different scenarios for string splitting.
	// Each test case includes a name to identify the specific scenario, the input string,
	// the separator used to split the string, the expected before part, the expected after part,
	// and whether the separator is expected to be found in the input string.
	cases := []struct {
		name           string
		input          string
		sep            string
		expectedBefore string
		expectedAfter  string
		expectedFound  bool
	}{
		{name: "separator in middle", input: "hello,world", sep: ",", expectedBefore: "hello", expectedAfter: "world", expectedFound: true},
		{name: "separator at beginning", input: ",world", sep: ",", expectedBefore: "", expectedAfter: "world", expectedFound: true},
		{name: "separator at end", input: "hello,", sep: ",", expectedBefore: "hello", expectedAfter: "", expectedFound: true},
		{name: "no separator", input: "helloworld", sep: ",", expectedBefore: "helloworld", expectedAfter: "", expectedFound: false},
		{name: "empty input string", input: "", sep: ",", expectedBefore: "", expectedAfter: "", expectedFound: false},
		{name: "empty separator", input: "helloworld", sep: "", expectedBefore: "helloworld", expectedAfter: "", expectedFound: false},
		{name: "long separator", input: "helloXXworld", sep: "XX", expectedBefore: "hello", expectedAfter: "world", expectedFound: true},
		{name: "separator not found in complex string", input: "abcdefg", sep: "123", expectedBefore: "abcdefg", expectedAfter: "", expectedFound: false},
	}

	// Iterate over each test case defined in the cases slice.
	// Each test case is executed within this loop, allowing the results to be independently validated.
	for _, tt := range cases {
		// Execute each test case as a subtest using t.Run, which provides a descriptive name for each test case.
		// This makes it easier to identify the results of individual tests when they are run.
		t.Run(tt.name, func(t *testing.T) {
			// Call the SplitStringBySeparator function with the input string and separator from the current test case.
			// The function will return the before part of the string, the after part, and whether the separator was found.
			before, after, found := SplitStringBySeparator(tt.input, tt.sep)
			// Assert that the returned before part matches the expected value from the test case.
			// This ensures that the function correctly identifies the part of the string that precedes the separator.
			assert.Equal(t, tt.expectedBefore, before, "Before value mismatch")
			// Assert that the returned after part matches the expected value from the test case.
			// This checks whether the function correctly captures the part of the string that follows the separator.
			assert.Equal(t, tt.expectedAfter, after, "After value mismatch")
			// Assert that the found flag matches the expected value, indicating whether the separator was present in the input string.
			// This confirms that the function correctly identifies if the separator was found during the split operation.
			assert.Equal(t, tt.expectedFound, found, "Found flag mismatch")
		})
	}
}

// TestStringIsEmpty verifies the behavior of the StringIsEmpty function.
// This test checks if the function accurately determines whether a given string
// is empty or consists solely of whitespace characters. It covers multiple scenarios,
// including completely empty strings, strings made up of only whitespace characters,
// and normal strings. By asserting the output against expected results, the test
// ensures that the StringIsEmpty function performs as intended across different inputs.
func TestStringIsEmpty(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", true},
		{"Whitespace Only", "   ", true},
		{"Newline Only", "\n", true},
		{"Tab Only", "\t", true},
		{"Normal String", "Hello", false},
		{"String with Whitespace", " Hello ", false},
		{"Multiple Whitespaces", "   Hello   ", false},
		{"Mixed Spaces and Newline", "   \n   ", true},
	}

	// Iterate through each test case defined in the cases slice.
	// This loop executes each test case and validates the results independently.
	for _, tt := range cases {
		// Execute each test case as a subtest using t.Run, providing a descriptive name for each test.
		// This allows for clearer identification of results for each scenario when tests are run.
		t.Run(tt.name, func(t *testing.T) {
			// Call the StringIsEmpty function with the input string from the current test case.
			// This function checks whether the input string is empty or consists solely of whitespace.
			result := StringIsEmpty(tt.input)
			// Assert that the result from StringIsEmpty matches the expected value from the test case.
			// The assert.Equal function checks if the actual result equals the expected value,
			// confirming that the function behaves as intended for this specific input.
			assert.Equal(t, tt.expected, result, "Expected result mismatch for input: %s", tt.input)
		})
	}
}
