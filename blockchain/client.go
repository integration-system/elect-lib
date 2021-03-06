package blockchain

import (
	"fmt"
	"github.com/integration-system/elect-lib/blockchain/internal"
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"sync"
)

const (
	maxAuthRetries = 10
)

var (
	json = jsoniter.ConfigFastest
)

type Client interface {
	RegisterVotersList(req RegisterVoterListRequest) (*RegisterVotersListResponse, error)
	IssueBallot(req IssueBallotRequest) (*IssueBallotResponse, error)
	RevokeParticipation(req VoterRequest) (*RevokeParticipationResponse, error)
	StoreBallot(req []byte) (*StoreBallotResponse, error)
}

type client struct {
	transport     internal.Transport
	cfg           BchConfig
	headers       map[string]string
	authenticated bool
	mx            sync.RWMutex
}

func NewClient(config BchConfig) Client {
	return &client{
		cfg:       config,
		headers:   map[string]string{"Content-Type": "application/json"},
		transport: internal.NewHttpTransport(),
	}
}

func (b *client) RegisterVotersList(req RegisterVoterListRequest) (*RegisterVotersListResponse, error) {
	request, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(RegisterVotersListResponse)
	err = b.invoke(registerVotersList, request, response, 0)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *client) RevokeParticipation(req VoterRequest) (*RevokeParticipationResponse, error) {
	request, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(RevokeParticipationResponse)
	err = b.invoke(revokeParticipation, request, response, 0)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *client) IssueBallot(req IssueBallotRequest) (*IssueBallotResponse, error) {
	request, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(IssueBallotResponse)
	err = b.invoke(issueBallot, request, response, 0)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *client) StoreBallot(req []byte) (*StoreBallotResponse, error) {
	response := new(StoreBallotResponse)
	err := b.invoke(storeBallot, req, response, 0)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *client) invoke(method string, request []byte, responsePtr interface{}, depth int) error {
	b.mx.RLock()
	authDone := b.authenticated
	b.mx.RUnlock()
	if !authDone {
		err := b.doAuth()
		if err != nil {
			return err
		}
	}

	response := Response{Result: responsePtr}
	statusCode, err := b.transport.Invoke(b.uri(method), b.headers, request, &response)
	if err != nil {
		return err
	}
	if statusCode == fasthttp.StatusUnauthorized {
		b.mx.Lock()
		b.authenticated = false
		b.mx.Unlock()
		if depth < maxAuthRetries {
			return b.invoke(method, request, responsePtr, depth+1)
		}
	}
	if statusCode >= fasthttp.StatusMultipleChoices {
		if response.Error != nil {
			return response.Error
		} else {
			return fmt.Errorf("unknown response: %v with status code: %d", response, statusCode)
		}
	}
	return nil
}

func (b *client) doAuth() error {
	b.mx.Lock()
	defer b.mx.Unlock()

	if b.authenticated {
		return nil
	}
	request := authenticateRequest{
		Login:    b.cfg.Login.Login,
		Password: b.cfg.Login.Password,
	}
	req, err := json.Marshal(request)
	if err != nil {
		return err
	}
	result := authenticateResponse{}
	response := Response{Result: &result}
	delete(b.headers, "Authorization")
	_, err = b.transport.Invoke(b.uri(authenticateMethod), b.headers, req, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return response.Error
	}

	b.authenticated = true
	b.headers["Authorization"] = "Bearer " + result.Token
	return nil
}

func (b *client) uri(method string) string {
	return b.cfg.Address + method
}
