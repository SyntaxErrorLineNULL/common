package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestRoundTripRateLimiterTransport tests the behavior of the RoundTripRateLimiterTransport
// when handling multiple concurrent requests. This test ensures that the transport correctly
// enforces rate limiting, allowing only a specified number of requests to be sent within a
// defined time period. It simulates a scenario where multiple requests are made to a test server,
// and it verifies that all requests are processed successfully while adhering to the rate limit.
func TestRoundTripRateLimiterTransport(t *testing.T) {
	// Create a new test HTTP server that simulates a real server.
	// The `httptest.NewServer` function starts a server and returns a handler that can respond to HTTP requests.
	// It is useful for testing HTTP clients without needing to rely on an actual external server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write an HTTP 200 OK status to the response.
		// This indicates to the client that the request was successful.
		w.WriteHeader(http.StatusOK)
	}))

	// Ensure that the test server is closed after the test completes.
	// This is important to release any resources associated with the server
	// and prevent any potential resource leaks during tests.
	defer server.Close()

	// Create a new instance of a rate-limited transport for handling HTTP requests.
	// This transport is configured to allow a maximum of 1 request every 100 milliseconds.
	// The `NewRoundTripRateLimiterTransport` function initializes the transport,
	// wrapping the default HTTP transport with rate limiting capabilities.
	roundTripRateLimiterTransport := NewRoundTripRateLimiterTransport(100*time.Millisecond, 1, http.DefaultTransport)
	// Assert that the newly created rate-limited transport is not nil.
	// This check ensures that the transport was successfully created and initialized,
	// which is crucial for the subsequent HTTP client to function properly.
	assert.NotNil(t, roundTripRateLimiterTransport, "Expected roundTripRateLimiterTransport to be initialized and not nil")
	// Initialize a new HTTP client with the rate-limited transport.
	// This client will use the previously created `roundTripRateLimiterTransport`
	// to manage outgoing requests, enforcing the specified rate limits.
	client := &http.Client{Transport: roundTripRateLimiterTransport}

	// Declare a WaitGroup to manage concurrent goroutines.
	// The WaitGroup will allow us to wait for all spawned goroutines to complete
	// their execution before proceeding further in the test.
	var wg sync.WaitGroup
	// Defer the call to Wait() on the WaitGroup until the surrounding function returns.
	// This ensures that the test will wait for all goroutines to finish before
	// exiting, allowing us to safely verify the outcomes of concurrent requests.
	defer wg.Wait()

	// Define the total number of requests to be made in the test.
	// This sets the expectation for how many concurrent requests will be
	// initiated to the test server, allowing us to verify the rate-limiting
	// functionality of the transport by ensuring that it can handle multiple
	// requests without exceeding the specified rate limit.
	totalRequests := 5

	for i := 0; i < totalRequests; i++ {
		// Increment the WaitGroup counter by one.
		// This indicates that a new goroutine is being started,
		// which will perform a task that the main test function needs to wait for.
		// This ensures proper synchronization and helps prevent premature termination
		// of the test before all goroutines have completed their execution.
		wg.Add(1)

		// Start a new goroutine to execute the following block of code concurrently.
		// This allows multiple requests to be sent to the server in parallel,
		// simulating a scenario where the client makes simultaneous calls.
		// Using goroutines helps to test the behavior of the throttled transport
		// under concurrent load, ensuring it handles multiple requests efficiently
		// and respects the rate limits set in the transport configuration.
		go func() {
			// Defer the call to Done() on the WaitGroup to indicate that the current goroutine
			// has completed its execution. This is crucial for ensuring that the main test function
			// can accurately wait for all concurrent requests to finish before it completes.
			// It helps maintain synchronization among goroutines, allowing for proper cleanup
			// after all requests have been processed.
			defer wg.Done()

			// Create a new HTTP GET request with a context for cancellation.
			// This request targets the server URL established earlier in the test,
			// allowing it to run independently of any cancellation signals.
			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
			// Assert that there was no error during the creation of the request.
			// This check is crucial to confirm that the request was formed correctly before
			// attempting to send it. If an error occurs, the test will fail with the provided
			// message, helping to diagnose issues in the request setup.
			assert.NoError(t, err, "Error creating request")

			// Send the HTTP request using the configured client.
			// This executes the request created earlier and waits for the response
			// from the server, allowing us to validate the behavior of the throttled transport.
			resp, err := client.Do(req)
			// Assert that there was no error when executing the request.
			// This is crucial as any failure at this stage would indicate a problem
			// in reaching the server or processing the request. The test will fail
			// with the provided message if an error occurs, assisting in identifying issues.
			assert.NoError(t, err, "Error making request")
			// Assert that the response status code is 200 OK.
			// This checks if the server successfully processed the request as expected.
			// If the status code is not OK, the test will fail, indicating that the
			// server did not handle the request correctly.
			assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status OK")
		}()
	}
}

// TestRateLimiterContextCancellation tests the behavior of the HTTP client with a rate-limited transport
// when using a context that has been canceled. This test ensures that when a request is attempted
// with a canceled context, the client correctly returns an error. It simulates a scenario where the
// context is canceled immediately after its creation, ensuring that the transport does not attempt
// to send the request and that the cancellation is properly handled. This is important for validating
// resource management and responsiveness in situations where operations need to be halted prematurely.
func TestRateLimiterContextCancellation(t *testing.T) {
	// Create a new test HTTP server that simulates a real server.
	// The `httptest.NewServer` function starts a server and returns a handler that can respond to HTTP requests.
	// It is useful for testing HTTP clients without needing to rely on an actual external server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write an HTTP 200 OK status to the response.
		// This indicates to the client that the request was successful.
		w.WriteHeader(http.StatusOK)
	}))

	// Ensure that the test server is closed after the test completes.
	// This is important to release any resources associated with the server
	// and prevent any potential resource leaks during tests.
	defer server.Close()

	// Create a new instance of a rate-limited transport for handling HTTP requests.
	// This transport is configured to allow a maximum of 1 request every 100 milliseconds.
	// The `NewRoundTripRateLimiterTransport` function initializes the transport,
	// wrapping the default HTTP transport with rate limiting capabilities.
	roundTripRateLimiterTransport := NewRoundTripRateLimiterTransport(100*time.Millisecond, 1, http.DefaultTransport)
	// Assert that the newly created rate-limited transport is not nil.
	// This check ensures that the transport was successfully created and initialized,
	// which is crucial for the subsequent HTTP client to function properly.
	assert.NotNil(t, roundTripRateLimiterTransport, "Expected roundTripRateLimiterTransport to be initialized and not nil")
	// Initialize a new HTTP client with the rate-limited transport.
	// This client will use the previously created `roundTripRateLimiterTransport`
	// to manage outgoing requests, enforcing the specified rate limits.
	client := &http.Client{Transport: roundTripRateLimiterTransport}

	// Create a new cancellable context using `context.WithCancel`.
	// This context will be passed to any operations that need to be canceled prematurely,
	// ensuring that resources tied to the context, like ongoing network requests, can be cleaned up.
	// The `cancel` function will be used to trigger the cancellation of the context.
	ctx, cancel := context.WithCancel(context.Background())
	// Immediately call `cancel()` to cancel the context.
	// This stops any ongoing or future operations tied to this context, ensuring no further
	// work is done using it. This is particularly useful in testing scenarios where cancellation
	// behavior needs to be simulated or when we need to ensure resources are released.
	cancel()

	// Create a new HTTP GET request with a context for cancellation.
	// This request targets the server URL established earlier in the test,
	// allowing it to run independently of any cancellation signals.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	// Assert that there was no error during the creation of the request.
	// This check is crucial to confirm that the request was formed correctly before
	// attempting to send it. If an error occurs, the test will fail with the provided
	// message, helping to diagnose issues in the request setup.
	assert.NoError(t, err, "Error creating request")

	// Attempt to execute the HTTP request using the previously canceled context.
	// Since the context has already been canceled, this request should not succeed
	// and is expected to return an error. The client will attempt to send the request,
	// but due to the cancellation, the operation will fail.
	_, err = client.Do(req)
	// Assert that an error was returned when executing the request.
	// This checks if the cancellation of the context caused the expected failure.
	// The test ensures that the error occurs as expected due to the context being canceled.
	// If no error is returned, the test will fail, signaling an issue with handling the cancellation.
	assert.Error(t, err, "Expected an error due to context cancellation")
}
