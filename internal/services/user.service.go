package services

import "example.com/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}




func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),


	}
}

func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}


