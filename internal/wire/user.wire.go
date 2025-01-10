//go:build wireinject

package wire

import (
	"example.com/go-ecommerce-backend-api/internal/controller"
	"example.com/go-ecommerce-backend-api/internal/repo"
	"example.com/go-ecommerce-backend-api/internal/services"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		services.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
