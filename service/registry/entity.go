package registry

import (
	"time"
)

const (
	VoterStateRevoked      VoterState = -1
	VoterStateEmpty        VoterState = 0
	VoterStateNotValid     VoterState = 1
	VoterStateValid        VoterState = 2
	VoterStateBallotIssued VoterState = 3
)

type VoterState int32

func (s VoterState) ToString() string {
	switch s {
	case VoterStateRevoked:
		return "Revoked"
	case VoterStateEmpty:
		return "Empty"
	case VoterStateNotValid:
		return "Invalid"
	case VoterStateValid:
		return "Valid"
	case VoterStateBallotIssued:
		return "BallotIssued"
	default:
		return "Unknown"
	}
}

func (s VoterState) CanVote() bool {
	switch s {
	case VoterStateValid:
		return true
	default:
		return false
	}
}

func (s VoterState) CanUpdateValidation() bool {
	switch s {
	case VoterStateNotValid, VoterStateValid:
		return true
	default:
		return false
	}
}

type (
	Voter struct {
		CitizenId             string `pg:",pk"`
		ElectionId            int    `pg:",pk"`
		State                 VoterState
		ValidationDescription map[string]interface{} `pg:",default:'',use_zero"`
		Extra                 map[string]interface{} `pg:",default:'',use_zero"`
		RequestId             string
		UpdatedAt             time.Time `pg:",default:'NOW()',use_zero"`
	}

	Citizen struct {
		Id        string `pg:",pk"`
		Version   int64
		Data      *Data
		UpdatedAt time.Time
	}

	Data struct {
		FirstName  string    `json:"first_name"`
		MiddleName string    `json:"middle_name"`
		LastName   string    `json:"last_name"`
		BirthDate  time.Time `json:"birth_date"`
		Sso        string    `json:"sso"`
		Passport   *Passport `json:"doc_passport_rf"`
		Phone      *Contact  `json:"contact_mobile_registration"`
		Snils      *Document `json:"doc_snils"`
		Address    *Address  `json:"addr_registration"`
	}

	Passport struct {
		Document
		Series string `json:"series"`
	}

	Document struct {
		Value      string `json:"value"`
		Deleted    bool   `json:"deleted"`
		Validation bool   `json:"validation"`
	}

	Address struct {
		Unom       string `json:"unom"`
		Unad       string `json:"unad"`
		Deleted    bool   `json:"deleted"`
		Validation bool   `json:"validation"`
	}

	Contact struct {
		Value   string `json:"value"`
		Deleted bool   `json:"deleted"`
	}
)
