package controller

import (
	"net/http"

	"example.com/go-ecommerce-backend-api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	// name := c.DefaultQuery("name", "anonymous")
	// // c.shouldBindJSON()
	// uid := c.Query("uid")
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	// 	"name":    name,
	// 	"uid":     uid,
	// 	"users":   []string{"user1", "user2"},
	// })

	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"user1", "user2"},
	})

}
