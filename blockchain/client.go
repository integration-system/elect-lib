package blockchain

import (
	"github.com/integration-system/elect-lib/blockchain/internal"
	json "github.com/json-iterator/go"
)

type Client interface {
	RegisterVotersList(req RegisterVoterListRequest) (*RegisterVotersListResponse, error)
	IssueBallot(req IssueBallotRequest) (*IssueBallotResponse, error)
	RegisterVoter(req RegisterVoterRequest) (*RegisterVoterResponse, error)
	StoreBallot(req []byte) (*StoreBallotResponse, error)
}

type client struct {
	cli *internal.Client
}

func NewBlockchainClient(config BlockchainConfig) *client {
	return &client{cli: internal.NewClient(config.Address, config.Login.Login, config.Login.Password)}
}

func (b *client) RegisterVotersList(req RegisterVoterListRequest) (*RegisterVotersListResponse, error) {
	request, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(RegisterVotersListResponse)
	err = b.cli.Invoke(registerVotersList, request, response)
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
	err = b.cli.Invoke(issueBallot, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *client) RegisterVoter(req RegisterVoterRequest) (*RegisterVoterResponse, error) {
	request, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(RegisterVoterResponse)
	err = b.cli.Invoke(registerVoter, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *client) StoreBallot(req []byte) (*StoreBallotResponse, error) {
	response := new(StoreBallotResponse)
	err := b.cli.Invoke(storeBallot, req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
