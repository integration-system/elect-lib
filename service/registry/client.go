package registry

import (
	"github.com/integration-system/isp-lib/v2/backend"
)

type Service interface {
	DeleteVotersByElectionId(DeleteVotersByElectionRequest) (*VotersDeletedResponse, error)
}

type grpcClient struct {
	cc       *backend.RxGrpcClient
	callerId int
}

func (c *grpcClient) DeleteVotersByElectionId(req DeleteVotersByElectionRequest) (*VotersDeletedResponse, error) {
	res := new(VotersDeletedResponse)
	err := c.cc.Invoke(deleteVotersByElectionId, c.callerId, req, res)
	return res, err
}

func NewGrpcClient(cli *backend.RxGrpcClient, callerId int) Service {
	return &grpcClient{
		cc:       cli,
		callerId: callerId,
	}
}
