package account

import (
	"example.com/go-ecommerce-backend-api/internal/services"
	"example.com/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type cUserLogin struct {
}

var Login = new(cUserLogin)

func (c *cUserLogin) Login(ctx *gin.Context) {

	// implement logic for login
	err := services.UserLogin().Login(ctx)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErCodeSuccess, nil)

}
