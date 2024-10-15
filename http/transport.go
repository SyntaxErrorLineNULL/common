package http

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// RoundTripRateLimiterTransport is a custom HTTP transport that wraps an existing http.RoundTripper.
// It adds rate-limiting functionality to control the number of outgoing HTTP requests.
// This struct ensures that the transport honors rate-limiting constraints when making requests.
type RoundTripRateLimiterTransport struct {
	// roundTripperWrap is the underlying HTTP transport (an `http.RoundTripper`) that handles the actual HTTP request execution.
	// It can be any implementation of `http.RoundTripper`, such as the default transport or a custom one.
	roundTripperWrap http.RoundTripper

	// rateLimiter is a rate limiter that controls the rate of requests being sent through this transport.
	// It ensures that the transport doesn't exceed a predefined number of requests within a certain time period.
	// This is useful for APIs with rate-limiting policies, preventing overloading or being throttled by the server.
	rateLimiter *rate.Limiter
}

// NewRoundTripRateLimiterTransport creates a new RoundTripRateLimiterTransport to add rate-limiting to outgoing HTTP requests.
// It wraps an existing HTTP transport and limits the number of requests that can be made within a specified time period.
// This ensures that HTTP requests sent through this transport comply with rate-limiting rules to avoid overloading services.
func NewRoundTripRateLimiterTransport(limitPeriod time.Duration, requestCount int, transportWrap http.RoundTripper) http.RoundTripper {
	// Return a pointer to a newly created RoundTripRateLimiterTransport instance.
	// This transport will wrap the existing transportWrap and apply rate-limiting to outgoing requests.
	return &RoundTripRateLimiterTransport{
		roundTripperWrap: transportWrap,
		rateLimiter:      rate.NewLimiter(rate.Every(limitPeriod), requestCount),
	}
}

// RoundTrip executes a single HTTP request and returns its response.
// This method adheres to the http.RoundTripper interface, enabling RoundTripRateLimiterTransport
// to be used in conjunction with HTTP clients, allowing for rate-limited requests.
func (t *RoundTripRateLimiterTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	// Wait for permission from the rate limiter before proceeding with the request.
	// This line ensures that the request adheres to the specified rate limits by
	// blocking the execution until the rate limiter allows sending the request.
	err := t.rateLimiter.Wait(request.Context())

	// Check if there was an error while waiting for the rate limit to allow the request.
	// This can happen if the context is canceled (for example, if the request is
	// taking too long) or if there are issues with the rate limiting logic.
	if err != nil {
		// If an error occurred, return nil for the response along with the error.
		// This signals to the calling function that the request cannot proceed due
		// to exceeding the rate limit or context-related issues.
		// Returning nil response indicates that no valid HTTP response can be provided.
		return nil, err
	}

	// Call the wrapped transport's RoundTrip method to execute the actual HTTP request.
	// This sends the request to the specified URL and waits for a response from the server.
	// The response returned by this call will contain the status code, headers, and body
	// of the response, allowing the caller to handle it appropriately.
	return t.roundTripperWrap.RoundTrip(request)
}
