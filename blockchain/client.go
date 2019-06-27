package blockchain

import (
	"github.com/integration-system/isp-lib/http"
	"github.com/integration-system/isp-lib/logger"
)

type Blockchain struct {
	Login           BchLogin `valid:"required~Required"`
	Address         string   `valid:"required~Required"`
	VotingEventBody interface{}
}

type BchLogin struct {
	Login    string `valid:"required~Required"`
	Password string `valid:"required~Required"`
}

var (
	bchServiceClient = http.NewJsonRestClient()
	Client           = newBlockChain(bchServiceClient)
)

func ReceiveConfiguration(bch Blockchain) {
	Client.ReceiveConfiguration(bch)
	if _, err := Client.Authenticate(); err != nil {
		logger.Fatal(err)
	}
}
