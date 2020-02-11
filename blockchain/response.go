package blockchain

import "fmt"

type Response struct {
	Error *struct {
		Code    int
		Details []interface{}
		Message string
	}
	Result interface{}
}

func (b Response) ConvertError() error {
	if b.Error == nil {
		return nil
	}
	return BchError{
		error: fmt.Sprintf("code: '%d' message '%s' details: '%v'", b.Error.Code, b.Error.Message, b.Error.Details),
	}
}

type BchError struct {
	error string
}

func (b BchError) Error() string {
	return b.error
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
