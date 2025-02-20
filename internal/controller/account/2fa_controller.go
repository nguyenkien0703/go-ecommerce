package account

import (
	"example.com/go-ecommerce-backend-api/internal/model"
	"example.com/go-ecommerce-backend-api/internal/services"
	"example.com/go-ecommerce-backend-api/internal/utils/context"
	"example.com/go-ecommerce-backend-api/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type sUser2FA struct {
}

var TwoFA = new(sUser2FA)

// User Setup Two Factor Authentication
// @Summary      ser Setup Two Factor Authentication
// @Description  ser Setup Two Factor Authentication
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization token"
// @Param        payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /user/two-factor/setup [post]
func (c *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		// Handle error
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid setupTwoFactorAuth parameter")
		return
	}
	// get UserId from uuid (token)
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	fmt.Sprintln("ctx.Request.Context() line 36---", ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "UserId is not valid")
		return
	}
	log.Println("UserId: ", userId)
	params.UserId = uint32(userId)

	codeResult, err := services.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}
