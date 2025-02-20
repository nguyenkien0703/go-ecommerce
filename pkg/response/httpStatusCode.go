package response

const (
	ErrCodeSuccess      = 201 // success
	ErrCodeParamInvalid = 203 // email is invalid
	ErrInvalidToken     = 301 // token is invalid

	// register code
	ErrCodeUserHasExist = 501 // user has exist
	ErrInvalidOTP       = 302
	ErrSendEmailOTP     = 303

	// Err Login
	ErrCodeOtpNotExists     = 609
	ErrCodeUserOtpNotExists = 608

	// User Authentication
	ErrCodeAuthFailed = 405

	// Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed  = 801
	ErrCodeTwoFactorAuthVerifyFailed = 802
)

var msg = map[int]string{
	ErrCodeSuccess:          "success",
	ErrCodeParamInvalid:     "email is invalid",
	ErrInvalidToken:         "token is invalid",
	ErrCodeUserHasExist:     "user has exist",
	ErrInvalidOTP:           "otp error",
	ErrSendEmailOTP:         "send email otp error",
	ErrCodeOtpNotExists:     "OTP exists but not registered",
	ErrCodeUserOtpNotExists: "User OTP not exists",
	ErrCodeAuthFailed:       "Authentication failed",
	// Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed:  "Two Factor Authentication setup failed",
	ErrCodeTwoFactorAuthVerifyFailed: "Two Factor Authentication verify failed",
}
