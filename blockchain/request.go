package blockchain

type RegisterVoterRequest struct {
	VoterId string
}

type RegisterVoterListRequest struct {
	VotersIds []string
}

type IssueBallotRequest struct {
	VoterId  string
	VotingId int
	Version  string
	Fio      Fio
	Passport *Passport `json:",omitempty"`
}

type Fio struct {
	Firstname  string
	Surname    string
	Patronymic string
}

type Passport struct {
	Series string
	Number string
}
