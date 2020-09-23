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

func GetUser(ctx *gin.Context) {
	var req request.GetUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse user ID in GetUser[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.GetUser(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle GetUser req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func CreateUser(ctx *gin.Context) {
	var req request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse CreateUser req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckCreateUser(req); err != nil {
		log.Errorf("Failed to validate CreateUser req[%v]: %v", req, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	resp, err := service.CreateUser(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle CreateUser req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)

}

func UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse UpdateUser req[%v]: %v",ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckUpdateUser(req); err != nil {
		log.Errorf("Failed to validate UpdateUser req[%v]: %v", req, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	if err := service.UpdateUser(req); err != nil {
		msg := fmt.Sprintf("Failed to handle UpdateUser req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func DeleteUser(ctx *gin.Context) {
	var req request.DeleteUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse user ID in DeleteUser[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := service.DeleteUser(req); err != nil {
		msg := fmt.Sprintf("Failed to handle DeleteUser req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
