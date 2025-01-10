package services

import (
	"example.com/go-ecommerce-backend-api/internal/repo"
	"example.com/go-ecommerce-backend-api/pkg/response"
)

//type UserService struct {
//	userRepo *repo.UserRepo
//}
//
//func NewUserService() *UserService {
//	return &UserService{
//		userRepo: repo.NewUserRepo(),
//
//
//	}
//}
//
//func (us *UserService) GetInfoUser() string {
//	return us.userRepo.GetInfoUser()
//}
//
//

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	// 1 check email exist
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExist
	}

	return response.ErCodeSuccess

}
