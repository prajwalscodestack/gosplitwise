package models

const (
	JWT_KEY = "mV8LW#9yFmN3!a@2tq$6JsPZ"
)

const (
	STATUS_SUCCESS = "SUCCESS"
	STATUS_FAILED  = "FAILED"
)

type Response struct {
	Error   string
	Message string
	Result  interface{}
	Status  string
}
type Credential struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
