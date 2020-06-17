package blockchain

import "github.com/integration-system/elect-lib/blockchain/transport"

type Option func(c *client)

func WithTransport(t transport.Transport) Option {
	return func(c *client) {
		c.transport = t
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(c *client) {
		for k, v := range headers {
			c.headers[k] = v
		}
	}
}
