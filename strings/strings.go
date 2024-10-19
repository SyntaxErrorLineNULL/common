package strings

import "strings"

// SplitStringBySeparator takes an input string and a separator, then splits the input string into two parts:
// before the separator and after the separator, while also indicating whether the separator was found.
// If the separator is found, it returns the part of the string before the separator, the part after it, and true.
// If the separator is not found, it returns the original string as before, an empty string as after, and false.
func SplitStringBySeparator(input, sep string) (before, after string, found bool) {
	// Calculate the length of the separator for later use.
	sepLen := len(sep)

	// Check if the length of the separator is zero, which indicates that an empty
	// separator has been provided. This is important because splitting by an empty
	// string does not yield meaningful results and should be handled explicitly.
	if sepLen == 0 {
		// Return the original input string as the before result, an empty string
		// as the after result, and false to indicate that no valid separator was found.
		return input, "", false
	}

	// Find the index of the first occurrence of the separator in the input string.
	// The strings.Index function returns the index of the separator or -1 if it's not found.
	if i := strings.Index(input, sep); i >= 0 {
		// If the separator is found (i >= 0), split the string into two parts:
		// input[:i] gives the substring before the separator.
		// input[i+len(sep):] gives the substring after the separator.
		// Return true indicating that the separator was found.
		return input[:i], input[i+sepLen:], true
	}

	// If the separator was not found (i < 0), return the original input as before,
	// an empty string as after, and false to indicate the separator wasn't found.
	return input, "", false
}
