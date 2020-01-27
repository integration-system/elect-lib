package internal

import (
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

type Client struct {
	login            string
	password         string
	headers          map[string]string
	completeInitAuth bool
	cli              *fasthttp.HostClient
	mx               sync.Mutex
}

func NewClient(address, login, password string) *Client {
	cli := &Client{
		login:    login,
		password: password,
		headers:  map[string]string{"Content-Type": "application/json"},
		cli:      &fasthttp.HostClient{Addr: address},
	}
	return cli
}

func (b *Client) Invoke(url string, request []byte, responsePtr interface{}) error {
	response := Response{Result: responsePtr}
	statusCode, err := b.invokeWithConvertResponse(url, request, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		if statusCode == fasthttp.StatusUnauthorized {
			err := b.auth()
			if err == nil {
				_, err = b.invokeWithConvertResponse(url, request, &response)
				if err != nil {
					return err
				}
				if response.Error != nil {
					return response.ConvertError()
				}
				return nil
			}
		}
		return response.ConvertError()
	}
	return nil
}

func (b *Client) invokeWithConvertResponse(url string, request []byte, respPtr interface{}) (int, error) {
	response, statusCode, err := b.invoke(url, request)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(response, respPtr)
	if err != nil {
		return 0, err
	}
	return statusCode, nil
}

func (b *Client) invoke(uri string, request []byte) ([]byte, int, error) {
	if !b.completeInitAuth {
		err := b.auth()
		if err != nil {
			return nil, 0, err
		}
	}
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(uri)
	req.Header.SetMethod(fasthttp.MethodPost)
	if len(b.headers) > 0 {
		for key, value := range b.headers {
			req.Header.Set(key, value)
		}
	}
	if request != nil {
		req.SetBody(request)
	}

	err := b.cli.DoTimeout(req, res, time.Second*15)
	if err != nil {
		return nil, 0, err
	}

	statusCode := res.StatusCode()
	response := res.Body()
	return response, statusCode, nil
}

func (b *Client) auth() error {
	b.mx.Lock()
	b.completeInitAuth = true
	b.mx.Unlock()
	request := authenticateRequest{
		Login:    b.login,
		Password: b.password,
	}
	req, err := json.Marshal(request)
	if err != nil {
		return err
	}
	result := authenticateResponse{}
	response := Response{Result: &result}
	_, err = b.invokeWithConvertResponse(authenticateMethod, req, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return response.ConvertError()
	}
	b.mx.Lock()
	b.headers["Authorization"] = "Bearer " + result.Token
	b.mx.Unlock()
	return nil
}
