package response

const (
	ErCodeSuccess       = 201   // success
	ErrCodeParamInvalid = 203 // email is invalid
)

var msg = map[int]string{
	ErCodeSuccess:       "success",
	ErrCodeParamInvalid: "email is invalid",
}
