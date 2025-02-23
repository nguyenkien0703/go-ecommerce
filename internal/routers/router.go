package routers

//
//import (
//	"fmt"
//
//	c "example.com/go-ecommerce-backend-api/internal/controller"
//	"example.com/go-ecommerce-backend-api/internal/middlewares"
//	"github.com/gin-gonic/gin"
//)
//
//func AA() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Println("before -> AA")
//		c.Next()
//		fmt.Println("alter -> AA")
//	}
//}
//func BB() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Println("before -> BB")
//		c.Next()
//		fmt.Println("alter -> BB")
//	}
//}
//
//func CC(c *gin.Context) {
//	fmt.Println("before -> CC")
//	c.Next()
//	fmt.Println("alter -> CC")
//
//}
//
//func NewRouter() *gin.Engine {
//	r := gin.Default()
//	// use middleware
//	r.Use(middlewares.AuthenMiddleware(), BB(), CC)
//
//	v1 := r.Group("/v1")
//	{
//		v1.GET("/ping", c.NewPongController().Pong)
//		//v1.GET("/user/1", c.NewUserController().GetUserById)
//		// v1.PUT("/ping", Pong)
//		// v1.PATCH("/ping", Pong)
//		// v1.DELETE("/ping", Pong)
//		// v1.HEAD("/ping", Pong)
//		// v1.OPTIONS("/ping", Pong)
//	}
//
//	return r
//}
