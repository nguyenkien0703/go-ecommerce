package account

import (
	"example.com/go-ecommerce-backend-api/global"
	"example.com/go-ecommerce-backend-api/internal/model"
	"example.com/go-ecommerce-backend-api/internal/services"
	"example.com/go-ecommerce-backend-api/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	fmt.Println("params-------", params)
	codeStatus, err := services.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErCodeSuccess, nil)

}
