package user

import (
	"example.com/go-ecommerce-backend-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.PUT("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")

	}
}
