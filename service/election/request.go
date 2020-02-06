package election

type (
	SearchElectionRequest struct {
		Limit  int
		Offset int
		Filter FilterSearch
	}

	FilterSearch struct {
		Name string
	}

	IdentityRequest struct {
		Id int `valid:"required~Required"`
	}

	IdentityListRequest struct {
		IdList []int
	}
)
