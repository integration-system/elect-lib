package registry

import (
	"time"
)

const AllStates = 100

// Request

type GetAvailableElectionsRequest struct {
	SsoId string `valid:"required~Required"`
}

type CheckElectionRequest struct {
	SsoId      string `valid:"required~Required"`
	ElectionId int    `valid:"required~Required"`
}

type DebugScriptRequest struct {
	Data   *Data  `valid:"required~Required"`
	Script string `valid:"required~Required"`
	Extra  map[string]interface{}
}

type CanIssueBallotRequest struct {
	SsoId      string `valid:"required~Required"`
	ElectionId int    `valid:"required~Required"`
}

type DoIssueBallotRequest struct {
	SsoId      string `valid:"required~Required"`
	ElectionId int    `valid:"required~Required"`
}

type SetVoterStateRequest struct {
	CitizenId  string     `valid:"required~Required"`
	ElectionId int        `valid:"required~Required"`
	State      VoterState `valid:"in(-1|1|2|3)"`
}

type DeleteVotersByElectionRequest struct {
	ElectionId int `valid:"required~Required"`
}

type GetElectionStatsRequest struct {
	ElectionId int `valid:"required~Required"`
}

type ExportVotersRequest struct {
	ElectionId int `valid:"required~Required"`
	Filter     SearchVotersRequestFilter
}

type SearchVotersRequest struct {
	Offset     int
	Limit      int `valid:"range(1|10000)"`
	ElectionId int `valid:"required~Required"`
	Filter     SearchVotersRequestFilter
}

type SearchVotersRequestFilter struct {
	ByState    int
	BySnils    string
	ByPassport *struct {
		Series string `valid:"required~Required"`
		Number string `valid:"required~Required"`
	}
	ByFio *struct {
		FirstName  string `valid:"required~Required"`
		LastName   string `valid:"required~Required"`
		MiddleName string
		BirthDate  time.Time
	}
}

// Response

type CheckElectionResponse struct {
	Valid                 *bool
	ValidationDescription map[string]interface{}
	Extra                 map[string]interface{}
}

type DebugScriptResponse struct {
	Response *CheckElectionResponse
	Error    string
}

type IssueBallotResponse struct {
	Valid bool
	Voter
}

type SearchVotersResponse struct {
	TotalCount int
	Items      []CitizenVoter
}

type VotersDeletedResponse struct {
	Count int
}

type ElectionStatsResponse struct {
	RevokedCount      int
	InvalidCount      int
	ValidCount        int
	BallotIssuedCount int
	TotalCount        int
}

type CitizenVoter struct {
	Voter
	Data *Data
}
