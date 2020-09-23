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

func CreateFavorite(ctx *gin.Context) {
	var req request.CreateFavoriteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse CreateFavorite req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.CreateFavorite(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle CreateOrder req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func ListFavorite(ctx *gin.Context) {
	var req request.ListFavoriteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse ListFavorite req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckListFavorite(&req); err != nil {
		log.Errorf("Failed to validate ListFavorite req[%v]: %v", req, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
	}

	resp, err := service.ListFavorite(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle ListFavorite req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func DeleteFavorite(ctx *gin.Context) {
	var req request.DeleteFavoriteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse DeleteFavorite req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	err := service.DeleteFavorite(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to DeleteFavoriteby req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
