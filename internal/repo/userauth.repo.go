package repo

import (
	"example.com/go-ecommerce-backend-api/global"
	"fmt"
	"time"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email) // usr:email:otp
	fmt.Sprint("key---- " + key)
	fmt.Sprint("otp---- " + string(otp))

	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
