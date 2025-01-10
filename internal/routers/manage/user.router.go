package manage

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	//ur := repo.NewUserRepository()
	//us := services.NewUserService(ur)
	//userHandlerNonDependency := controller.NewUserController(us)
	//userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.POST("/register")
		//userRouterPublic.PUT("/otp")

	}

	// private router
	userRouterPrivate := Router.Group("/admin/user")
	//userRouterPrivate.Use(limiter()) //hạn chế số lượng yêu cầu từ một IP nhất định
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")

	}
}
