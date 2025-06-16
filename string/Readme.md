# String Manipulation Utilities in Go

 The string package provides utility functions for string manipulation in Go. It includes functions for splitting strings based on separators or width constraints, checking for empty or whitespace-only strings, and formatting strings with specific case rules. The package is designed to handle Unicode characters correctly and includes robust error handling for edge cases.

### Installation

To use this package in your Go project, ensure you have Go installed, then import it into your code. The package depends on the standard library packages strings and unicode/utf8.

# Functions

## `SplitStringBySeparator`

This function splits an input string into two parts based on the first occurrence of a separator. It returns the part of the string before the separator, the part after the separator, and a boolean indicating if the separator was found.

### Signature
```go
func SplitStringBySeparator(input, sep string) (before, after string, found bool)
```

### Parameters:
- `input`: The string to be split.
- `sep`: The separator string to split on.

### Returns:
- `before`: The part of the string before the separator.
- `after`: The part of the string after the separator.
- `found`: bool

### Example:
```go
before, after, found := SplitStringBySeparator("apple,banana", ",")
fmt.Println(before) // Output: "apple"
fmt.Println(after)  // Output: "banana"
fmt.Println(found)  // Output: true
```

## `IsEmpty`

This function checks whether a string is empty or contains only whitespace characters. It returns `true` if the string is empty or consists solely of whitespace, and `false` otherwise.

### Signature
```go
func IsEmpty(str string) bool
```

### Parameter:
- `str`: The string to check.

### Returns:
- `bool`: true if the string is empty or contains only whitespace. False otherwise.

### Example:
```go
isEmpty := IsEmpty("   ")
fmt.Println(isEmpty) // Output: true
```

## `SplitStringWithWidthConstraints`

This function splits a string into multiple segments based on the specified width constraints (`maxWidth` and `overflowWidth`). It ensures no words are broken across segments.

### Signature
```go
func SplitStringWithWidthConstraints(str string, maxWidth, overflowWidth int) []string
```

### Parameters:
- `str`: The input string to split.
- `maxWidth`: The maximum width for each segment.
- `overflowWidth`: Additional width allowed for overflow.

### Returns:
- A slice of strings where:
    - Each string represents a segment of the input string fitting within the width constraints.
    - Words are kept intact without being broken across segments.

### Example:
```go
segments := SplitStringWithWidthConstraints("This is a sample string for testing width constraints", 10, 5)
fmt.Println(segments) // Output: ["This is a", "sample", "string for", "testing width", "constraints"]
```

## `UpperCaseFirst`

This function converts the first non-whitespace character of a string to uppercase, while converting the rest of the string to lowercase. It ignores leading whitespace and returns the original string unchanged if it's empty or consists only of whitespace.

### Signature
```go
func UpperCaseFirst(str string) string
```

### Parameter:
- `str`: The input string.

### Returns:
- The string with:
    - The first non-whitespace character converted to uppercase.
    - The rest of the string converted to lowercase.
    - The original string unchanged if it contains only whitespace.

### Example:
```go
formatted := UpperCaseFirst(" hello WORLD")
fmt.Println(formatted) // Output: "Hello world"
```
