package domain

import (
	"gitlab.alx/mdm-ext/mdm-elect-lib/blockchain"
	"gitlab.alx/mdm-ext/mdm-elect-lib/digest"
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
