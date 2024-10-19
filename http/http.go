package http

import (
	"context"
	"errors"
	"net/http"
	"sync"
)

type Client struct {
	wg         *sync.WaitGroup
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		wg:         &sync.WaitGroup{},
		httpClient: &http.Client{},
	}
}

func (c *Client) Invoke(ctx context.Context, request *Request) (*http.Response, error) {
	if request == nil {
		return nil, errors.New("empty request")
	}

	newRequest, err := c.buildRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(newRequest)
}

func (c *Client) buildRequest(ctx context.Context, request *Request) (*http.Request, error) {
	body := request.Body
	if body == nil {
		body = http.NoBody
	}

	newRequest, err := http.NewRequestWithContext(ctx, request.Method, request.URL.String(), body)
	if err != nil {
		return nil, err
	}

	if request.Header != nil {
		for key, values := range *request.Header {
			for _, value := range values {
				newRequest.Header.Add(key, value)
			}
		}
	}

	if len(request.Cookies) > 0 {
		for _, cookie := range request.Cookies {
			newRequest.AddCookie(cookie)
		}
	}

	return newRequest, nil
}
