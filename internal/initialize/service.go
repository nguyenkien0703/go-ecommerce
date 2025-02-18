package initialize

import (
	"example.com/go-ecommerce-backend-api/global"
	"example.com/go-ecommerce-backend-api/internal/database"
	"example.com/go-ecommerce-backend-api/internal/services"
	"example.com/go-ecommerce-backend-api/internal/services/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// user service interface.....
	services.InitUserLogin(impl.NewUserLoginImpl(queries))
}
