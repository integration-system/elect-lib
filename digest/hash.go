package digest

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const (
	randomLength = 1000000000
)

type MessageDigest struct {
	Timestamp  int64  `json:"timestamp"`
	Random     int64  `json:"random"`
	SecureHash string `json:"secureHash"`
}

func MakeDigest(secret string) MessageDigest {
	timestamp := time.Now().Unix()
	random := rand.Int63n(randomLength)

	return MessageDigest{
		Timestamp:  timestamp,
		Random:     random,
		SecureHash: makeDigest(timestamp, random, secret),
	}
}

func makeDigest(timestamp int64, random int64, secret string) string {
	digestStr := fmt.Sprintf("%d|%d|%s", timestamp, random, secret)
	f := sha256.New()
	f.Write([]byte(digestStr))
	return hex.EncodeToString(f.Sum(nil))
}
