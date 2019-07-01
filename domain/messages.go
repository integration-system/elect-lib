package domain

import (
	"time"
)

type MessageDigest struct {
	Timestamp  int64  `json:"timestamp"`
	Random     int64  `json:"random"`
	SecureHash string `json:"secureHash"`
}

type ElectMessage struct {
	MessageDigest
	Id   string
	Time time.Time
	Type string
}

type RegisterVoterMessage struct {
	ElectMessage
	Payload string
}
