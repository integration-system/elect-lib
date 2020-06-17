package transport

import "github.com/valyala/fasthttp"

type Option func(transport *httpTransport)

func WithClient(c *fasthttp.Client) Option {
	return func(transport *httpTransport) {
		transport.cli = c
	}
}
