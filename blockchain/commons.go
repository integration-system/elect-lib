package blockchain

import (
	"github.com/integration-system/isp-lib/http"
	"gitlab.alx/mdm-ext/mdm-elect-lib/blockchain/domain"
)

type blockChainRep struct {
	bch     Blockchain
	headers map[string]string
	client  http.RestClient
}

func newBlockChain(client http.RestClient) *blockChainRep {
	return &blockChainRep{client: client}
}

func (b *blockChainRep) ReceiveConfiguration(bch Blockchain) {
	b.headers = map[string]string{"Content-Type": "application/json"}
	b.bch = bch
}

func (b *blockChainRep) Authenticate() (*domain.LoginResponse, error) {
	request := domain.AuthenticateRequest{
		Login:    b.bch.Login.Login,
		Password: b.bch.Login.Password,
	}
	result := domain.LoginResponse{}
	response := domain.Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+authenticate, b.headers, request, &response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	b.headers["Authorization"] = "Bearer " + result.Token
	return &result, nil
}

func (b *blockChainRep) Flush() error {
	response := domain.Response{}
	if err := http.NewJsonRestClient().Invoke("DELETE", b.bch.Address+flush, b.headers, nil, &response); err != nil {
		return err
	} else if response.Error != nil {
		return response.ConvertError()
	}
	return nil
}

func (b *blockChainRep) CreateVotingEvent() (*domain.VotingEventResponse, error) {
	result := domain.VotingEventResponse{}
	response := domain.Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+createVotingEvent, b.headers, b.bch.VotingEventBody, &response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	return &result, nil
}

func (b *blockChainRep) RegisterVotersList(req domain.RegisterVoterRequest) (*domain.RegisterVotersListResponse, error) {
	result := domain.RegisterVotersListResponse{}
	response := domain.Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+registerVotersList, b.headers, req, response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	return &result, nil
}

func (b *blockChainRep) IssueBallot(req domain.IssueBallotRequest) (*domain.Response, error) {
	result := domain.IssueBallotResponse{}
	response := domain.Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+issueBallot, b.headers, req, &response); err != nil {
		return &response, err
	} else {
		return &response, nil
	}
}

func (b *blockChainRep) RegisterVoter(req domain.RegisterVoterRequest) (*domain.RegisterVoterResponse, error) {
	result := domain.RegisterVoterResponse{}
	response := domain.Response{Result: &result}
	if err := b.client.Invoke("POST", b.bch.Address+registerVoter, b.headers, req, &response); err != nil {
		return nil, err
	} else if response.Error != nil {
		return nil, response.ConvertError()
	}
	return &result, nil
}
