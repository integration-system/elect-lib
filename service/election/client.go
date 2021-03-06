package election

import (
	"github.com/integration-system/isp-lib/v2/backend"
)

type grpcClient struct {
	cc       *backend.RxGrpcClient
	callerId int
}

func (c *grpcClient) GetCurrentRegElectionsByIdList(req IdentityListRequest) ([]Election, error) {
	res := make([]Election, 0)
	err := c.cc.Invoke(getCurrentRegElectionsByIdList, c.callerId, req, &res)
	return res, err
}

func (c *grpcClient) GetCurrentVotingElectionsByIdList(req IdentityListRequest) ([]Election, error) {
	res := make([]Election, 0)
	err := c.cc.Invoke(getCurrentVotingElectionsByIdList, c.callerId, req, &res)
	return res, err
}

func (c *grpcClient) GetElectionById(req IdentityRequest) (*Election, error) {
	res := new(Election)
	err := c.cc.Invoke(getElectionsById, c.callerId, req, res)
	return res, err
}

func (c *grpcClient) GetRegElectionById(req IdentityRequest) (*Election, error) {
	res := new(Election)
	err := c.cc.Invoke(getRegElectionById, c.callerId, req, res)
	return res, err
}

func (c *grpcClient) GetVotingElectionById(req IdentityRequest) (*Election, error) {
	res := new(Election)
	err := c.cc.Invoke(getVotingElectionById, c.callerId, req, res)
	return res, err
}

func NewGrpcClient(cli *backend.RxGrpcClient, callerId int) Service {
	return &grpcClient{
		cc:       cli,
		callerId: callerId,
	}
}
