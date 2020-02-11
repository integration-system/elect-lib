package election

import "time"

type (
	Election struct {
		Id int `valid:"required~Required"`
		ElectionInfo
	}

	ElectionInfo struct {
		Name               string `valid:"required~Required"`
		Description        string
		Extra              map[string]interface{}
		VotingStart        time.Time `valid:"required~Required"`
		VotingEnd          time.Time `valid:"required~Required"`
		RegStart           time.Time `valid:"required~Required"`
		RegEnd             time.Time `valid:"required~Required"`
		EtalonVotingStatus string    `valid:"in(BY_DATE|ENABLE|DISABLE)~Required in (BY_DATE|ENABLE|DISABLE)"`
		EtalonRegStatus    string    `valid:"in(BY_DATE|ENABLE|DISABLE)~Required in (BY_DATE|ENABLE|DISABLE)"`
		Filter             ElectionFilter
		CreatedAt          time.Time
	}

	ElectionFilter struct {
		Script string
	}
)
