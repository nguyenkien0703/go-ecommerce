package main

import (
	_ "example.com/go-ecommerce-backend-api/cmd/swag/docs"
	"example.com/go-ecommerce-backend-api/internal/initialize" // name of module, in go.mod file
	swaggerFiles "github.com/swaggo/files"                     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"                 // gin-swagger middleware
)

// @title           API Documentation Ecommerce Backend SHOPDEVGO
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/nguyenkien0703/go-ecommerce

// @contact.name   TEAM TIPSGO
// @contact.url    github.com/nguyenkien0703/go-ecommerce
// @contact.email  nguyenkien07032003ns@gmail.com

// @license.name  Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2024
// @schema http
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	// r := routers.NewRouter()
	// r.Run()// default 8080, if you want to change port, you can pass it as an argumnent to Run(":8081")
	r := initialize.Run()
	// use swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")

}
