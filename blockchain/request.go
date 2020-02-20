package blockchain

type VoterRequest struct {
	VoterId string
}

type RegisterVoterListRequest struct {
	VotersIds []string
}

type IssueBallotRequest struct {
	VoterId    string
	DistrictId int
	Version    string
}

type authenticateRequest struct {
	Login    string
	Password string
}

type authenticateResponse struct {
	Login string
	Token string
}
