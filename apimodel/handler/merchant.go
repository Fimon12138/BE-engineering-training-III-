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

func GetMerchant(ctx *gin.Context) {
	var req request.GetMerchantRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse merchant ID in GetMerchant[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.GetMerchant(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle GetMerchant req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func CreateMerchant(ctx *gin.Context) {
	var request request.CreateMerchantRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		msg := fmt.Sprintf("Failed to parse CreateMerchant req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckCreateMerchant(request); err != nil {
		log.Errorf("Failed to validate CreateMerchant req[%v]: %v", request, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	resp, err := service.CreateMerchant(request)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle CreateMerchant req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func UpdateMerchant(ctx *gin.Context) {
	var request request.UpdateMerchantRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		msg := fmt.Sprintf("Failed to parse UpdateMerchant req[%v]: %v", &ctx, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckUpdateMerchant(request); err != nil {
		log.Errorf("Failed to validate UpdateMerchant req[%v]: %v", request, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	if err := service.UpdateMerchant(request); err != nil {
		msg := fmt.Sprintf("Failed to handle UpdateMerchant req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func DeleteMerchant(ctx *gin.Context) {
	var req request.DeleteMerchantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse DeleteMerchant req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := service.DeleteMerchant(req); err != nil {
		msg := fmt.Sprintf("Failed to handle DeleteMerchant req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
