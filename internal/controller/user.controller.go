package controller

import (
	"example.com/go-ecommerce-backend-api/internal/services"
	"example.com/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

//type UserController struct {
//	userService *services.UserService
//}
//
//func NewUserController() *UserController {
//	return &UserController{
//		userService: services.NewUserService(),
//	}
//}
//
//func (uc *UserController) GetUserById(c *gin.Context) {
//
//	// response.SuccessResponse(c, 201, []string{"user1", "user2"})
//
//	response.ErrorResponse(c, 203, "no need!!!")
//
//}

type UserController struct {
	userService services.IUserService
}

func NewUserController(
	userService services.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)

}
