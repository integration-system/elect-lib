package domain

import (
	"github.com/integration-system/elect-lib/blockchain"
	"github.com/integration-system/elect-lib/digest"
	"time"
)

type ElectMessage struct {
	digest.MessageDigest
	Id   string
	Time time.Time
	Type string
}

type RegisterVoterMessage struct {
	ElectMessage
	Payload string
}

type IssueBallotMessage struct {
	ElectMessage
	Payload blockchain.IssueBallotRequest
}
