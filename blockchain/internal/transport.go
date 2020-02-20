package internal

import (
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	json = jsoniter.ConfigFastest
)

type Transport interface {
	Invoke(url string, headers map[string]string, request []byte, responsePtr interface{}) (int, error)
}

type httpTransport struct {
	cli *fasthttp.HostClient
}

func (b *httpTransport) Invoke(uri string, headers map[string]string, request []byte, respPtr interface{}) (int, error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(uri)
	req.Header.SetMethod(fasthttp.MethodPost)
	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	if request != nil {
		req.SetBody(request)
	}

	err := b.cli.DoTimeout(req, res, time.Second*15)
	if err != nil {
		return 0, err
	}

	statusCode := res.StatusCode()
	if respPtr != nil {
		err = json.Unmarshal(res.Body(), respPtr)
		if err != nil {
			return 0, err
		}
	}
	return statusCode, nil
}

func NewHttpTransport(address string) Transport {
	cli := &httpTransport{
		cli: &fasthttp.HostClient{Addr: address},
	}
	return cli
}
