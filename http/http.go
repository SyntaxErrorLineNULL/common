package http

import (
	"context"
	"sync"
)

type Client struct {
	wg *sync.WaitGroup
}

func NewClient() *Client {
	return &Client{wg: &sync.WaitGroup{}}
}

func (c *Client) Invoke(ctx context.Context) {}
