package manage

import (
	"example.com/go-ecommerce-backend-api/internal/controller"
	"example.com/go-ecommerce-backend-api/internal/repo"
	"example.com/go-ecommerce-backend-api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	ur := repo.NewUserRepository()
	us := services.NewUserService(ur)
	userHandlerNonDependency := controller.NewUserController(us)
	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.POST("/register", userHandlerNonDependency.Register)
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
