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
	Fio      Fio
	Passport Passport
}

type AuthenticateRequest struct {
	Login    string
	Password string
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
