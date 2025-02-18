package response

const (
	ErCodeSuccess       = 201 // success
	ErrCodeParamInvalid = 203 // email is invalid
	ErrInvalidToken     = 301 // token is invalid

	// register code
	ErrCodeUserHasExist = 501 // user has exist
	ErrInvalidOTP       = 302
	ErrSendEmailOTP     = 303

	// Err Login
	ErrCodeOtpNotExists = 609
)

var msg = map[int]string{
	ErCodeSuccess:       "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrInvalidToken:     "token is invalid",
	ErrCodeUserHasExist: "user has exist",
	ErrInvalidOTP:       "otp error",
	ErrSendEmailOTP:     "send email otp error",
	ErrCodeOtpNotExists: "OTP exists but not registered",
}
