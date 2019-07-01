package digest

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gitlab.alx/mdm-ext/mdm-elect-lib/domain"
	"math/rand"
	"time"
)

const (
	randomLength = 1000000000
)

func MakeDigest(secret string) domain.MessageDigest {
	timestamp := time.Now().Unix()
	random := rand.Int63n(randomLength)

	return domain.MessageDigest{
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
