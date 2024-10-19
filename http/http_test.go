package http

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func mockServer(t *testing.T) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with status OK and echo back the request method and headers using fmt.Sprintf
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, fmt.Sprintf("Method: %s; Headers: %s", r.Method, r.Header.Get("Test-Header")))
	})

	return httptest.NewServer(handler)
}

func TestClientInvoke(t *testing.T) {
	// Start a mock server
	server := mockServer(t)
	defer server.Close()

	// Create a new request with a valid URL and method
	req := &Request{
		RequestID: "test-123",
		Method:    http.MethodGet,
		URL:       &url.URL{Host: server.URL},
		Link:      server.URL,
		Header:    &http.Header{},
		Body:      bytes.NewBufferString("test body"),
	}

	// Set a header
	req.SetHeaders(map[string]string{"Test-Header": "TestValue"})

	// Create a new client
	client := NewClient()

	// Invoke the request
	response, err := client.Invoke(context.Background(), req)
	assert.NoError(t, err)

	defer response.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Check the response body
	var buf bytes.Buffer
	io.Copy(&buf, response.Body)
	expected := "Method: GET; Headers: TestValue"
	assert.Contains(t, buf.String(), expected)
}
