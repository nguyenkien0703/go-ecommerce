package services

import (
	"example.com/go-ecommerce-backend-api/internal/repo"
	"example.com/go-ecommerce-backend-api/internal/utils/crypto"
	"example.com/go-ecommerce-backend-api/internal/utils/random"
	"example.com/go-ecommerce-backend-api/internal/utils/sendto"
	"example.com/go-ecommerce-backend-api/pkg/response"
	"fmt"
	"strconv"
	"time"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {

	//0. hashEmail
	hashEmail := crypto.GetHash(email)
	fmt.Println("hashEmail: %s", hashEmail)

	//5. check otp is available
	//6. user spam otp
	//1. check email is exist db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExist
	}

	//2.new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST" {
		otp = 123456
	}

	//3.save OTP in redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	// 4. send Email OTP
	err = sendto.SendTemplateEmailOtp([]string{email}, "anonystick@gmail.com", "otp-auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})
	fmt.Printf("err sendto :::%d\n", err)
	if err != nil {
		return response.ErrSendEmailOTP
	}

	return response.ErCodeSuccess

}
