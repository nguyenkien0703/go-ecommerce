package routers

import (
	"example.com/go-ecommerce-backend-api/internal/routers/manage"
	"example.com/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
