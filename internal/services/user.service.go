package services

import (
	"context"
	"encoding/json"
	"example.com/go-ecommerce-backend-api/global"
	"example.com/go-ecommerce-backend-api/internal/repo"
	"example.com/go-ecommerce-backend-api/internal/utils/crypto"
	"example.com/go-ecommerce-backend-api/internal/utils/random"
	"example.com/go-ecommerce-backend-api/pkg/response"
	"fmt"
	"github.com/segmentio/kafka-go"
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
	fmt.Sprintln("otp----", otp)

	////3.save OTP in redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	//// 4. send Email OTP
	//err = sendto.SendTemplateEmailOtp([]string{email}, "anonystick@gmail.com", "otp-auth.html", map[string]interface{}{
	//	"otp": strconv.Itoa(otp),
	//})
	//fmt.Printf("err sendto :::%d\n", err)
	//if err != nil {
	//	return response.ErrSendEmailOTP
	//}

	// send email OTP by java
	//err = sendto.SendEmailToJavaByAPI(strconv.Itoa(otp), email, "opt-auth.html")
	//fmt.Printf("err sendto JAVA:::%d\n", err)
	//if err != nil {
	//	return response.ErrSendEmailOTP
	//}

	// send OTP via kafka JAVA

	body := make(map[string]interface{})
	body["otp"] = otp
	body["email"] = email
	bodyRequest, errK := json.Marshal(body)
	if errK != nil {
		fmt.Println(errK)
		return 1
	}
	fmt.Sprintln("string(bodyRequest)-----", string(bodyRequest))

	message := kafka.Message{
		Key:   []byte("otp-auth"),
		Value: []byte(string(bodyRequest)),
		Time:  time.Now(),
	}
	err = global.KafkaProducer.WriteMessages(context.Background(), message)
	if err != nil {
		return response.ErrSendEmailOTP
	}

	return response.ErrCodeSuccess

}
