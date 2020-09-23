package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tickethub_service/apimodel/request"
	"tickethub_service/apimodel/validate"
	"tickethub_service/service"
	"tickethub_service/util/errors"
	"tickethub_service/util/log"
)

func GetAccount(ctx *gin.Context) {
	var req request.GetAccountRequest

	if err :=ctx.ShouldBindJSON(&req); req.AccountName == "" {
		msg := fmt.Sprintf("Failed to parse account name in GetAccount[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.GetAccount(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to get account by req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func CreateAccount(ctx *gin.Context) {
	var req request.CreateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse Createaccount req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckCreateAccount(req); err != nil {
		log.Errorf("Failed to validate CreateAccount req[%v]: %v", req, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
	}

	resp, err := service.CreateAccount(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle CreateAccount req[%v] : %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func UpdateAccount(ctx *gin.Context) {
	var req request.UpdateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse UpdateAccount req:[%v]", ctx.Request)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := service.UpdateAccount(req); err != nil {
		msg := fmt.Sprintf("Failed to handle UpdateAccount req[%v]: %v", req, &err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func DeleteAccount(ctx *gin.Context) {
	var req request.DeleteAccountRequest
	if  err := ctx.ShouldBindJSON(&req); err != nil  {
		msg := fmt.Sprintf("Failed to parse account name in DeleteAccount[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := service.DeleteAccount(req); err != nil {
		msg := fmt.Sprintf("Failed to delete account by name[%v] : %v", ctx, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func LogIn(ctx *gin.Context) {
	var req request.LogIn
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse login req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.LogIn(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to  login req:[%v]", ctx.Request)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)

}

func SignUp(ctx *gin.Context) {
	var req request.SignUp
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to  parse signup req:[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckSignup(req); err != nil {
		log.Errorf("Failed to validate req[%v]: &v", err, req)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	if err := service.SignUp(req); err != nil {
		msg := fmt.Sprintf("Failed to  handle signup req:[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
	}

	ctx.AbortWithStatus(http.StatusOK)
}
