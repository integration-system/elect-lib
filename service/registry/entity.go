package registry

type (
	Voter struct {
		CitizenId             int
		ElectionId            int
		ElectStatus           string
		ValidationDescription interface{}
		Extra                 interface{}
		RegistrationLocked    bool
	}

	Citizen struct {
		Id      int
		Version int64
		Data    interface{}
	}
)
