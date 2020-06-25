package transport

import (
	"github.com/valyala/fasthttp"
	"time"
)

type Option func(transport *httpTransport)

func WithClient(c *fasthttp.Client) Option {
	return func(transport *httpTransport) {
		transport.cli = c
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(transport *httpTransport) {
		transport.timeout = timeout
	}
}
