package http

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSetHeaders verifies the functionality of the SetHeaders method in the Request struct.
// The test ensures that the method correctly adds and updates headers in various scenarios.
// It uses a set of predefined cases that cover multiple situations, such as setting new headers,
// replacing existing ones, and handling cases where no headers are provided. By asserting the
// final state of headers against the expected outcomes, the test confirms that the SetHeaders
// method behaves as intended and maintains the integrity of the headers within the Request object.
func TestSetHeaders(t *testing.T) {
	// Define a slice of test cases for the SetHeaders method. Each test case includes a descriptive name
	// to identify the scenario being tested, the initial state of headers in the Request object, a map
	// representing headers that should be added or updated, and the expected state of headers after
	// invoking the SetHeaders method. This structure allows for easy iteration over various conditions
	// to ensure the SetHeaders method functions correctly across different inputs and expected outcomes.
	cases := []struct {
		name            string
		initialHeaders  *http.Header
		headersToSet    map[string]string
		expectedHeaders *http.Header
	}{
		{
			name:           "Set multiple headers",
			initialHeaders: &http.Header{},
			headersToSet: map[string]string{
				"Content-Type": "application/json",
				"User-Agent":   "GoTest",
			},
			expectedHeaders: &http.Header{
				"Content-Type": []string{"application/json"},
				"User-Agent":   []string{"GoTest"},
			},
		},
		{
			name: "Replace existing header",
			initialHeaders: &http.Header{
				"Content-Type": []string{"text/plain"},
			},
			headersToSet: map[string]string{
				"Content-Type": "application/json",
			},
			expectedHeaders: &http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		{
			name:            "Set no headers",
			initialHeaders:  &http.Header{},
			headersToSet:    map[string]string{},
			expectedHeaders: &http.Header{},
		},
	}

	// Iterate over each test case defined in the `cases` slice.
	// Each test case is represented by `tt`, which contains the initial headers, headers to set,
	// and the expected headers after the SetHeaders method is invoked.
	// This loop allows for executing multiple scenarios, ensuring that the SetHeaders method behaves
	// correctly across a variety of input conditions and expected outcomes.
	for _, tt := range cases {
		// Each test case in `cases` is run within this loop.
		// The `t.Run` function is used to run subtests, which allows for the execution of multiple
		// test scenarios independently while providing descriptive names for each case.
		// The purpose of this loop is to verify the behavior of the SetHeaders method for different sets
		// of initial and additional headers. It ensures that the method functions correctly across various scenarios.
		t.Run(tt.name, func(t *testing.T) {
			// The test begins by initializing a Request object with predefined initial headers.
			// Then, it invokes the SetHeaders method with additional headers to be set or updated.
			// Finally, it asserts that the final headers in the Request object match the expected headers,
			// ensuring that the method behaves correctly in terms of setting and updating headers.
			req := &Request{Header: tt.initialHeaders}
			// Call the SetHeaders method on the Request object.
			// This will set the headers specified in `tt.headersToSet`,
			// which might replace or add to the existing headers in `req.Header`.
			req.SetHeaders(tt.headersToSet)
			// Use assert.Equal to compare the final headers in the Request object to the expected headers.
			// This assertion checks that the headers were updated correctly by the SetHeaders method.
			// If the headers do not match, the test will fail with a descriptive message,
			// indicating that the method's behavior did not align with the expected outcome.
			assert.Equal(t, tt.expectedHeaders, req.Header, "Headers did not match the expected result after SetHeaders call")
		})
	}
}

// TestSetMethod verifies the functionality of the SetMethod method in the Request struct.
// The test ensures that the method correctly sets valid HTTP methods and handles invalid inputs.
// It utilizes a set of predefined cases to cover various scenarios, including valid methods (GET, POST, DELETE)
// and an invalid method, which should result in an error. By asserting the final state of the Method field
// against the expected outcomes, the test confirms that the SetMethod method behaves as intended
// and enforces proper validation of HTTP methods within the Request object.
func TestSetMethod(t *testing.T) {
	// Define a slice of test cases for the SetMethod function.
	// Each test case is represented by an anonymous struct that includes
	// a name to identify the specific scenario being tested,
	// the HTTP method to be set on the Request object during the test,
	// a boolean indicating whether an error is expected when setting the method,
	// and the expected value of the Method field in the Request object after the method call.
	cases := []struct {
		name      string
		method    string
		expectErr bool
		expected  string
	}{
		{"Valid GET Method", "GET", false, "GET"},
		{"Valid POST Method", "POST", false, "POST"},
		{"Invalid Method", "INVALID", true, ""},
		{"Valid DELETE Method", "delete", false, "DELETE"},
	}

	// Iterate over each test case defined in the `cases` slice.
	// Each test case is represented by `tt`, which contains the method to set,
	// the expected outcome, and whether an error is anticipated.
	for _, tt := range cases {
		// Each test case in `cases` is executed within this loop.
		// The `t.Run` function allows the execution of subtests, providing
		// descriptive names for each test case, making it easier to identify results.
		t.Run(tt.name, func(t *testing.T) {
			// Create a new instance of the Request struct,
			// initializing its Header field with a pointer to a new http.Header object.
			// This sets up a fresh Request object that can be used for testing,
			// ensuring that there are no pre-existing headers that could affect the outcome.
			req := &Request{Header: &http.Header{}}
			// Call the SetMethod method on the Request object, passing in the HTTP method
			// specified in the current test case represented by `tt.method`.
			// The purpose of this call is to test whether the method can be set correctly
			// and to verify if any errors occur based on the validity of the provided method.
			err := req.SetMethod(tt.method)

			// Check if the current test case expects an error to occur.
			// This conditional evaluates the boolean field `expectErr` from the test case struct,
			// which indicates whether the invocation of the method should result in an error or not.
			if tt.expectErr {
				// If an error is expected, assert that an error was indeed returned by the method call.
				// The assert.Error function checks that the error variable is not nil,
				// which confirms that the method behaved as intended by signaling an invalid operation.
				assert.Error(t, err)
			} else {
				// If no error is expected, assert that the error variable is nil, indicating success.
				// The assert.NoError function checks that the error is nil,
				// ensuring that the method call completed without encountering issues.
				assert.NoError(t, err)
				// Verify that the Method field in the Request object matches the expected value.
				// The assert.Equal function checks if the value of `req.Method` equals `tt.expected`,
				// which was predefined in the test case, confirming that the method was set correctly.
				assert.Equal(t, tt.expected, req.Method)
			}
		})
	}
}

// TestSetURL validates the behavior of the SetURL method in the Request struct.
// This test ensures that the method correctly sets a valid URL without any errors
// and that it rejects a nil URL by returning a specific error. It verifies both
// valid and invalid scenarios, checking that the URL is properly assigned when
// valid and that an appropriate error is returned for invalid input.
func TestSetURL(t *testing.T) {
	// Initialize a new Request object with an empty HTTP header.
	// This Request object will be used to test the SetURL method, specifically verifying
	// that the URL can be set correctly without any issues. The empty header setup ensures
	// that there are no conflicts with the URL setting.
	req := &Request{Header: &http.Header{}}

	// Parse a valid URL string into a URL object.
	// This valid URL represents the new value to be assigned to the request.
	// We use the `url.Parse` function to convert the string into a URL object that can
	// be passed to the SetURL method for validation and assignment.
	validURL, _ := url.Parse("https://example.com")
	// Call the SetURL method to set the parsed URL in the Request.
	// This action updates the URL field of the Request object with the valid URL.
	// The method should not return any errors since the input URL is valid.
	err := req.SetURL(validURL)
	// Assert that no error was returned when setting the URL.
	// This ensures that the SetURL method handled the valid URL correctly,
	// indicating that the method functions as expected when provided with valid input.
	assert.NoError(t, err, "Expected no error for valid URL")
	// Assert that the request's URL matches the provided valid URL.
	// This verifies that the SetURL method correctly updated the request's URL field.
	// If the two URLs match, it confirms that the method successfully assigned the value.
	assert.Equal(t, validURL, req.URL, "Expected request URL to match the set URL")

	// Call the SetURL method with a nil URL, simulating invalid input.
	// Since the URL is nil, the method is expected to return an error,
	// signaling that a nil value is not allowed for the URL field.
	err = req.SetURL(nil)
	// Assert that an error was returned when trying to set a nil URL.
	// This ensures that the SetURL method properly detects invalid input
	// and prevents a nil value from being assigned to the request's URL.
	assert.Error(t, err, "Expected error for nil URL")
	// Assert that the error message matches the expected value.
	// This verifies that the method returns a specific error message
	// indicating that the URL cannot be nil, ensuring clear error reporting.
	assert.Equal(t, "URL cannot be nil", err.Error(), "Expected specific error message for nil URL")
}
