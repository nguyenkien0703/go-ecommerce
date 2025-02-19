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

// UpdatePasswordRegister
// @Summary      UpdatePasswordRegister
// @Description  UpdatePasswordRegister
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdatePasswordRegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /user/update_pass_register [post]
func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params model.UpdatePasswordRegisterInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	result, err := services.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)
	if err != nil {
		response.ErrorResponse(ctx, result, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErCodeSuccess, result)

}

// Verify OTP Login By User
// @Summary      Verify OTP Login By User
// @Description  Verify OTP Login By User
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body model.VerifyInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /user/verify_account [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	result, err := services.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErCodeSuccess, result)

}

// User Login
// @Summary      User Login
// @Description  User Login
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body model.LoginInputHaha true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /user/login [post]
func (c *cUserLogin) Login(ctx *gin.Context) {

	// implement logic for login
	var params model.LoginInputHaha
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeRs, dataRs, err := services.UserLogin().Login(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, dataRs)

}

// Register godoc
// @Summary      Register a new account
// @Description  When user register, system will send OTP to user's phone number or email address
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /user/register [post]
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
