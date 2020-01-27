package internal

const (
	authenticateErrorCode = 401
	authenticateMethod    = "/authenticate"
)

type authenticateRequest struct {
	Login    string
	Password string
}

type authenticateResponse struct {
	Login string
	Token string
}
