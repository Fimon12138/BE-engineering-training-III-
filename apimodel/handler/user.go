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
	var request request.GetUserRequest

	if request.ID = ctx.Param("ID"); request.ID == "" {
		msg := fmt.Sprintf("Failed to parse user ID in GetUser:[%v]", ctx.Request)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.GetUser(request)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle GetUser req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func CreateUser(ctx *gin.Context) {
	var request request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		msg := fmt.Sprintf("Failed to parse CreateUser req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckCreateUser(request); err != nil {
		log.Errorf("Failed to validate CreateUser req[%v]: %v", request, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	resp, err := service.CreateUser(request)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle CreateUser req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)

}

func UpdateUser(ctx *gin.Context) {
	var request request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		msg := fmt.Sprintf("Failed to parse UpdateUser req[%v]: %v",ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckUpdateUser(request); err != nil {
		log.Errorf("Failed to validate UpdateUser req[%v]: %v", request, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	if err := service.UpdateUser(request); err != nil {
		msg := fmt.Sprintf("Failed to handle UpdateUser req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func DeleteUser(ctx *gin.Context) {
	var request request.DeleteUserRequest
	if request.ID = ctx.Param("ID"); request.ID == "" {
		msg := fmt.Sprintf("Failed to parse user ID in DeleteUser:[%v]", ctx.Request)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := service.DeleteUser(request); err != nil {
		msg := fmt.Sprintf("Failed to handle DeleteUser req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
