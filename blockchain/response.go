package blockchain

import "fmt"

type BchError struct {
	Code    int
	Details []interface{}
	Message string
}

func (err *BchError) Error() string {
	return fmt.Sprintf("code: '%d' message '%s' details: '%v'", err.Code, err.Message, err.Details)
}

type Response struct {
	Error  *BchError
	Result interface{}
}

type VotingEventResponse struct {
	VotersRegistryAddress  string
	BallotsRegistryAddress string
}

type IssueBallotResponse struct {
	IssuedFor issuedFor
}

type issuedFor struct {
	VoterId  string
	VotingId int
}

type RegisterVoterResponse struct {
	VoterId string
	Added   bool
}

type RegisterVotersListResponse struct {
	VotersAdded   int
	VotersExisted int
}

type StoreBallotResponse struct {
	Stored bool
}
