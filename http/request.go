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
