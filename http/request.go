package http

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	URL    *url.URL
	Method string
	Header http.Header
	Body   io.Reader
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
