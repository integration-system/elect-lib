package internal

import "fmt"

type Response struct {
	Error *struct {
		Code    int
		Details []string
		Message string
	}
	Result interface{}
}

type ErrorResponse struct {
	StatusCode int
	Status     string
	Body       string
}

func (r ErrorResponse) Error() string {
	return fmt.Sprintf("statusCode:%d  status:%s  body:%s", r.StatusCode, r.Status, r.Body)
}

func (b Response) ConvertError() error {
	if b.Error == nil {
		return nil
	}
	return BchError{
		error: fmt.Sprintf("Code: %d Message %s Detaild %s", b.Error.Code, b.Error.Message, b.Error.Details),
	}
}

type BchError struct {
	error string
}

func (b BchError) Error() string {
	return b.error
}
