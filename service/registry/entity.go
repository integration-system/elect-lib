package registry

type (
	Voter struct {
		CitizenId             int
		ElectionId            int
		ElectStatus           string
		ValidationDescription map[string]interface{}
		Extra                 map[string]interface{}
		RegistrationLocked    bool
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
