package registry

import (
	"time"
)

const (
	VoterStateRevoked  VoterState = -1
	VoterStateEmpty    VoterState = 0
	VoterStateNotValid VoterState = 1
	VoterStateValid    VoterState = 2
	VoterStateVoted    VoterState = 3
)

type VoterState int32

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
		ValidationDescription map[string]interface{}
		Extra                 map[string]interface{}
		RequestId             string
		UpdatedAt             time.Time
	}

	Citizen struct {
		Id      int
		Version int64
		Data    Data
	}

	Data struct {
		FirstName  string    `json:"first_name"`
		MiddleName string    `json:"middle_name"`
		LastName   string    `json:"last_name"`
		BirthDate  string    `json:"birth_date"`
		Sso        string    `json:"sso"`
		Passport   *Passport `json:"doc_passport_rf"`
		Phone      *Contact  `json:"contact_mobile_registration"`
		Snils      *Document `json:"doc_snils"`
		Address    *Address  `json:"addr_registration"`
	}

	Passport struct {
		Document
		Series string
	}

	Document struct {
		Value      string
		Deleted    bool
		Validation bool
	}

	Address struct {
		Unom       string
		Unad       string
		Deleted    bool
		Validation bool
	}

	Contact struct {
		Value   string
		Deleted bool
	}
)
