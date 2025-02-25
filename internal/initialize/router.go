package initialize

import (
	"example.com/go-ecommerce-backend-api/global"
	"example.com/go-ecommerce-backend-api/internal/middlewares"
	"example.com/go-ecommerce-backend-api/internal/routers"
	"github.com/gin-gonic/gin"
)

// Initial router
func InitRouter() *gin.Engine {
	// router := gin.Default()
	var router *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}
	// middlewares
	router.Use() // logger
	router.Use() // cross
	// middlewares limiter
	router.Use(middlewares.NewRateLimiter().GlobalLimiter())
	router.Use(middlewares.NewRateLimiter().PublicAPILimiter())
	router.Use(middlewares.NewRateLimiter().UserPrivateAPILimiter())
	// middlewares prometheus
	router.Use(middlewares.PrometheusMiddleware())

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User
	prometheusRouter := routers.RouterGroupApp.Prometheus
	MainGroup := router.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus") // tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
		//... other routes...
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
		//... other routes...
	}

	prometheusRouter.InitRouter(router)

	return router
}
