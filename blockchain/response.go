package blockchain

import "fmt"

type BchError struct {
	Code    int
	Details []interface{}
	Message string
}

func (err *BchError) Error() string {
	return fmt.Sprintf("code: '%d' message: '%s' details: '%v'", err.Code, err.Message, err.Details)
}

type Response struct {
	Error  *BchError
	Result interface{}
}

type IssueBallotResponse struct {
	IssuedFor IssuedFor
}

type IssuedFor struct {
	VoterId  string
	VotingId int
}

type RegisterVotersListResponse struct {
	VotersAdded   int
	VotersExisted int
}

type RevokeParticipationResponse struct {
	VoterId string
	Revoked bool
}

type StoreBallotResponse struct {
	Stored bool
}
