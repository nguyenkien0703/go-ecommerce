package initialize

import (
	"example.com/go-ecommerce-backend-api/global"
	"example.com/go-ecommerce-backend-api/pkg/logger"
)

// Initial Logger
func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
