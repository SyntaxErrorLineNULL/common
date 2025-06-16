package string

import (
	"strings"
	"unicode/utf8"
)

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

// IsEmpty checks if a given string is empty or contains only whitespace.
// It returns true if the string is empty or consists solely of whitespace characters,
// and false otherwise.
func IsEmpty(str string) bool {
	// Use strings.TrimSpace to remove leading and trailing whitespace from the string.
	// Check the length of the trimmed string. If the length is zero, it indicates that
	// the original string was either empty or contained only whitespace.
	return len(strings.TrimSpace(str)) == 0
}

// SplitStringWithWidthConstraints splits an input string into multiple segments
// based on specified width constraints: maxWidth and overflowWidth. Each segment
// adheres to the maxWidth limit while allowing overflowWidth, ensuring that no
// words are broken across segments. It returns a slice of strings, each representing
// a chunk of the original input string that fits within the defined width constraints.
func SplitStringWithWidthConstraints(str string, maxWidth, overflowWidth int) []string {
	// Check if maxWidth is less than 0, which would indicate an invalid negative value.
	// This ensures that maxWidth remains a valid non-negative value for further processing.
	if maxWidth < 0 {
		// Set maxWidth to 0 to handle invalid negative str, ensuring that the value used
		// for splitting the string is non-negative and won't cause unexpected behavior.
		maxWidth = 0
	}

	// Check if the number of runes (Unicode code points) in the str string is less than the sum
	// of maxWidth and overflowWidth. This condition ensures that the string is short enough to fit
	// within the allowed width without needing to be split.
	if utf8.RuneCountInString(str) < maxWidth+overflowWidth {
		// If the condition is true, return the str string as a single-element slice.
		// This avoids unnecessary processing when the string already fits within the allowed width.
		return []string{str}
	}

	// Create a 2D slice to hold chunks of words. The initial size is set to 1,
	// indicating that we will start with one chunk to store the words.
	chunks := make([][]string, 1)
	// Initialize the currentChunk variable to track the index of the chunk
	// that is currently being populated. This starts at 0, indicating the first chunk.
	currentChunk := 0
	// Initialize charCount to 0 to keep track of the total number of characters
	// (runes) added to the current chunk. This will help manage the width limits.
	charCount := 0

	// Split the str string into words using whitespace as the delimiter.
	// The strings.Fields function returns a slice of words, effectively
	// removing any leading or trailing whitespace from the str.
	words := strings.Fields(str)

	// Iterate over each word in the slice of words obtained from the str string.
	// The range keyword allows us to loop through the words slice, where
	// the variable word represents the current word in each iteration.
	// This loop processes each word individually, enabling us to manage
	// the chunking of the str string based on the defined width limits.
	for _, word := range words {
		// Calculate the number of runes (characters) in the current word
		// using utf8.RuneCountInString. This ensures we account for
		// multi-byte characters correctly when determining the word length.
		wordLength := utf8.RuneCountInString(word)

		// Check if adding the current word would exceed the maximum allowed width,
		// considering the overflow width. If it does exceed and the current chunk
		// is not empty, we need to start a new chunk for the next word.
		if charCount+wordLength > maxWidth+overflowWidth && len(chunks[currentChunk]) > 0 {
			// Move to the next chunk by incrementing the currentChunk index.
			// This allows us to begin filling the next chunk with new words.
			currentChunk++
			// Reset the character count to 0 for the new chunk,
			// as we are starting fresh with a new set of words.
			charCount = 0
			// Append a new empty slice to the chunks slice to represent the new chunk,
			// which will be filled with the next set of words.
			chunks = append(chunks, []string{})
		}

		// Add the current word to the current chunk's slice of words.
		// This appends the word to the slice located at the index currentChunk.
		chunks[currentChunk] = append(chunks[currentChunk], word)
		// Update the character count by adding the length of the current word.
		// This keeps track of how many characters are in the current chunk,
		// allowing us to manage the width constraints effectively.
		charCount += wordLength
	}

	// Create a new slice called result, initialized with zero length and a capacity
	// equal to the number of chunks. This pre-allocation optimizes memory usage
	// by allocating enough space to hold all the resulting strings from the chunking process.
	result := make([]string, 0, len(chunks))

	// Iterate over each chunk in the chunks slice.
	// The range keyword allows us to loop through the chunks slice, where
	// chunk represents the current chunk of words being processed in each iteration.
	for _, chunk := range chunks {
		// Join the words in the current chunk into a single string, separating them with spaces.
		// The strings.Join function concatenates the words, effectively reconstructing the
		// chunk as a single string, which is then appended to the result slice.
		result = append(result, strings.Join(chunk, " "))
	}

	// Return the final result slice, which contains the strings constructed
	// from the chunks of the str string based on the defined width limits.
	return result
}

// UpperCaseFirst takes a string as input and returns the same string
// with the first non-whitespace character converted to uppercase.
// If the input string is empty or consists only of whitespace, it returns
// the string unchanged. This function ensures that only the first character
// of the trimmed string is affected, while the rest of the characters are
// converted to lowercase, providing a standardized format for the output.
func UpperCaseFirst(str string) string {
	// Check if the input string is empty or contains only whitespace.
	// If it is, return the input string as-is, ensuring no changes are made.
	if IsEmpty(str) {
		return str
	}

	// Remove leading whitespace from the input string to focus on the first character.
	// This prepares the string for the uppercase conversion of the first character.
	trimmed := strings.TrimLeft(str, " ")

	// Convert the first character of the trimmed string to uppercase.
	// Use strings.ToUpper to change the case of the first character,
	// and then concatenate it with the rest of the trimmed string
	// converted to lowercase.
	return strings.ToUpper(trimmed[:1]) + strings.ToLower(trimmed[1:])
}
