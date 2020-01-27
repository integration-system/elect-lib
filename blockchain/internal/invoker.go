package internal

import (
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"time"
)

type Client struct {
	login    string
	password string
	headers  map[string]string
	cli      *fasthttp.HostClient
}

func NewClient(address, login, password string) *Client {
	cli := &Client{
		login:    login,
		password: password,
		headers:  map[string]string{"Content-Type": "application/json"},
		cli:      &fasthttp.HostClient{Addr: address},
	}
	_ = cli.auth()
	return cli
}

func (b *Client) Invoke(url string, request []byte, responsePtr interface{}) error {
	response := Response{Result: responsePtr}
	err := b.invokeWithConvertResponse(url, request, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		if response.Error.Code == authenticateErrorCode {
			err := b.auth()
			if err == nil {
				err = b.invokeWithConvertResponse(url, request, &response)
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

func (b *Client) invokeWithConvertResponse(url string, request []byte, respPtr interface{}) error {
	response, err := b.invoke(url, request)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response, respPtr)
	return err
}

func (b *Client) invoke(uri string, request []byte) ([]byte, error) {
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
		return nil, err
	}

	statusCode := res.StatusCode()
	if statusCode != fasthttp.StatusOK {
		return nil, ErrorResponse{
			StatusCode: statusCode,
			Status:     fasthttp.StatusMessage(statusCode),
			Body:       string(res.Body()),
		}
	}
	response := res.Body()
	return response, nil
}

func (b *Client) auth() error {
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
	err = b.invokeWithConvertResponse(authenticateMethod, req, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return response.ConvertError()
	}
	b.headers["Authorization"] = "Bearer " + result.Token
	return nil
}
