package string

import (
	"fmt"
	"strings"
	"testing"
)

// BenchmarkSplitStringBySeparator_LongString benchmarks the performance of the
// SplitStringBySeparator function with a long input string that contains a separator
// in the middle. This test evaluates how the function performs under conditions
// with significant input length and measures the execution time.
// a = 1000, b = 1000 		BenchmarkSplitStringBySeparator_LongString-12   106 513 381	        11.12 ns/op
// a = 1000000, b = 1000000 BenchmarkSplitStringBySeparator_LongString-12     107 406   	    9438 ns/op
func BenchmarkSplitStringBySeparator_LongString(b *testing.B) {
	// Create a long input string with 'a' repeated 1,000 times, followed by a comma
	// and then 'b' repeated 1,000 times. This serves as the test case for splitting.
	input := fmt.Sprintf("%s,%s", strings.Repeat("a", 1000), strings.Repeat("b", 1000))
	// Define the separator to be used for splitting the input string.
	separator := ","

	// Reset the timer to ensure that the benchmark only measures the execution time
	// of the SplitStringBySeparator function calls, excluding any setup time.
	b.ResetTimer()

	// Loop b.N times to execute the benchmark function multiple times.
	// This allows for accurate timing and performance evaluation of the function
	// across several iterations, providing a better average performance metric.
	for i := 0; i < b.N; i++ {
		// Call the SplitStringBySeparator function with the long input string
		// and the defined separator. This is the core operation being benchmarked,
		// measuring how quickly the function can split the long string at the separator.
		SplitStringBySeparator(input, separator)
	}
}
