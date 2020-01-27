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
	mx               sync.RWMutex
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
	b.mx.RLock()
	completeInitAuth := b.completeInitAuth
	b.mx.RUnlock()
	if !completeInitAuth {
		err := b.auth()
		if err != nil {
			return err
		}
	}

	response := Response{Result: responsePtr}
	statusCode, err := b.invoke(url, request, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		if statusCode == fasthttp.StatusUnauthorized {
			b.mx.Lock()
			b.completeInitAuth = false
			b.mx.Unlock()
			err := b.auth()
			if err == nil {
				_, err = b.invoke(url, request, &response)
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

func (b *Client) invoke(uri string, request []byte, respPtr interface{}) (int, error) {
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
		return 0, err
	}

	statusCode := res.StatusCode()
	err = json.Unmarshal(res.Body(), respPtr)
	if err != nil {
		return 0, err
	}
	return statusCode, nil
}

func (b *Client) auth() error {
	b.mx.Lock()
	defer b.mx.Unlock()

	if b.completeInitAuth {
		return nil
	}
	b.completeInitAuth = true
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
	_, err = b.invoke(authenticateMethod, req, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return response.ConvertError()
	}

	b.headers["Authorization"] = "Bearer " + result.Token
	return nil
}
