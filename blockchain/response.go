package blockchain

import "fmt"

type Response struct {
	Error *struct {
		Code    int
		Details []string
		Message string
	}
	Result interface{}
}

type LoginResponse struct {
	Login string
	Token string
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

type bchError struct {
	error string
}

func (b bchError) Error() string {
	return b.error
}

func (b Response) ConvertError() error {
	if b.Error == nil {
		return nil
	}
	return bchError{
		error: fmt.Sprintf("Code: %d Message %s Detaild %s", b.Error.Code, b.Error.Message, b.Error.Details),
	}
}
