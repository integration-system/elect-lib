package blockchain

import (
	"bytes"
	"github.com/integration-system/isp-lib/http"
	"io/ioutil"
	http2 "net/http"
)

type blockchainClient struct {
	bch     BlockchainConfig
	headers map[string]string
	client  http.RestClient
}

func NewBlockchainClient(client http.RestClient) *blockchainClient {
	return &blockchainClient{client: client}
}

func (b *blockchainClient) ReceiveConfiguration(bch BlockchainConfig) {
	b.headers = map[string]string{"Content-Type": "application/json"}
	b.bch = bch
}

func (b *blockchainClient) Authenticate() (*LoginResponse, error) {
	request := AuthenticateRequest{
		Login:    b.bch.Login.Login,
		Password: b.bch.Login.Password,
	}
	result := LoginResponse{}
	response := Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+authenticate, b.headers, request, &response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	b.headers["Authorization"] = "Bearer " + result.Token
	return &result, nil
}

func (b *blockchainClient) Flush() error {
	response := Response{}
	if err := http.NewJsonRestClient().Invoke("DELETE", b.bch.Address+flush, b.headers, nil, &response); err != nil {
		return err
	} else if response.Error != nil {
		return response.ConvertError()
	}
	return nil
}

func (b *blockchainClient) CreateVotingEvent() (*VotingEventResponse, error) {
	result := VotingEventResponse{}
	response := Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+createVotingEvent, b.headers, b.bch.VotingEventBody, &response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	return &result, nil
}

func (b *blockchainClient) RegisterVotersList(req RegisterVoterListRequest) (*RegisterVotersListResponse, error) {
	result := RegisterVotersListResponse{}
	response := Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+registerVotersList, b.headers, req, response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	return &result, nil
}

func (b *blockchainClient) IssueBallot(req IssueBallotRequest) (*IssueBallotResponse, error) {
	result := IssueBallotResponse{}
	response := Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+issueBallot, b.headers, req, &response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	return &result, nil
}

func (b *blockchainClient) RegisterVoter(req RegisterVoterRequest) (*RegisterVoterResponse, error) {
	result := RegisterVoterResponse{}
	response := Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+registerVoter, b.headers, req, &response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	return &result, nil
}

func (b *blockchainClient) StoreBallot(req []byte) ([]byte, error) {
	body := bytes.NewBuffer(req)
	request, err := http2.NewRequest(http2.MethodPost, b.bch.Address+storeBallot, body)
	if err != nil {
		return nil, err
	}
	defer func() { _ = request.Body.Close() }()

	for key, value := range b.headers {
		request.Header.Set(key, value)
	}

	response, err := http2.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()

	if resp, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
