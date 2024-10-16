package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	RequestID string
	URL       *url.URL
	Header    *http.Header
	Method    string
	Body      io.Reader
}

// SetHeaders is a method that allows setting multiple headers for an HTTP request.
// It accepts a map of string key-value pairs representing the header names and their respective values.
// The method iterates over the provided headers and sets each one in the request's Header field.
// The headers are applied using the standard `http.Header.Set()` method, which sets the header to the specified value,
// replacing any existing values associated with the given header name.
//
// This method returns the `*Request` itself, allowing for method chaining, so you can easily
// configure the request with additional headers or other parameters in a fluent style.
func (r *Request) SetHeaders(headers map[string]string) *Request {
	// Iterate over the map of headers where each key is the header name
	// and the corresponding value is the header value to be set.
	for key, value := range headers {
		// Set the header in the request using the http.Header.Set method.
		// This will overwrite any existing value for the same header name.
		r.Header.Set(key, value)
	}

	// Return the request object itself to allow method chaining.
	return r
}

// SetMethod sets the HTTP method for the Request object. It validates the provided method against
// a predefined list of acceptable HTTP methods and returns an error if the method is invalid.
// The input method is converted to uppercase to standardize its format.
func (r *Request) SetMethod(method string) error {
	// Convert the provided method string to uppercase to ensure consistent formatting.
	// This allows for case-insensitive method checks, as HTTP methods are conventionally represented in uppercase.
	method = strings.ToUpper(method)

	// Define a map of valid HTTP methods for validation.
	// The map keys represent HTTP methods, while the empty struct values serve as placeholders.
	validMethods := map[string]struct{}{
		http.MethodGet:     {},
		http.MethodPost:    {},
		http.MethodPut:     {},
		http.MethodDelete:  {},
		http.MethodPatch:   {},
		http.MethodOptions: {},
		http.MethodHead:    {},
	}

	// Check if the provided method exists in the validMethods map.
	// This ensures that only recognized HTTP methods are allowed, enhancing the robustness of the Request object.
	if _, exists := validMethods[method]; !exists {
		// If the method is not valid, return an error indicating that the specified method is invalid.
		// The error message includes the invalid method for better debugging and clarity.
		return errors.New(fmt.Sprintf("invalid HTTP method: %s", method))
	}

	// If the method is valid, set the Method field of the Request object to the validated method.
	// This updates the request to use the specified HTTP method for subsequent operations.
	r.Method = method

	// Return nil to indicate that the method was successfully set without any errors.
	// This allows the caller to proceed with confidence that the Request object is now in a valid state.
	return nil
}
