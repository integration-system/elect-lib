package blockchain

type Response struct {
	Error *struct {
		Code    int
		Details []string
		Message string
	}
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
