package user

import (
	"example.com/go-ecommerce-backend-api/internal/controller/account"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	//userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/user")
	{
		//userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/register", account.Login.Register)
		//userRouterPublic.PUT("/otp")
		userRouterPublic.POST("/login", account.Login.Login) // login -> YES -> No
		userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		userRouterPublic.POST("/update_pass_register", account.Login.UpdatePasswordRegister)

	}

	// private router
	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)

	}
}
