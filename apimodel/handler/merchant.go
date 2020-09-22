package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	request2 "tickethub_service/apimodel/request"
	validate2 "tickethub_service/apimodel/validate"
	service2 "tickethub_service/service"
	"tickethub_service/util/errors"
	"tickethub_service/util/log"
)

func GetMerchant(ctx *gin.Context) {
	var request request2.GetMerchantRequest

	if request.ID = ctx.Param("ID"); request.ID == "" {
		msg := fmt.Sprintf("Failed to parse merchant ID in GetMerchant:[%v]", ctx)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service2.GetMerchant(request)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle GetMerchant req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func CreateMerchant(ctx *gin.Context) {
	var request request2.CreateMerchantRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		msg := fmt.Sprintf("Failed to parse CreateMerchant req[%v]: %v", ctx, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate2.CheckCreateMerchant(request); err != nil {
		log.Errorf("Failed to validate CreateMerchant req[%v]: %v", request, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	resp, err := service2.CreateMerchant(request)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle CreateMerchant req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func UpdateMerchant(ctx *gin.Context) {
	var request request2.UpdateMerchantRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		msg := fmt.Sprintf("Failed to parse UpdateMerchant req[%v]: %v", ctx, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate2.CheckUpdateMerchant(request); err != nil {
		log.Errorf("Failed to validate UpdateMerchant req[%v]: %v", request, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	if err := service2.UpdateMerchant(request); err != nil {
		msg := fmt.Sprintf("Failed to handle UpdateMerchant req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func DeleteMerchant(ctx *gin.Context) {
	var request request2.DeleteMerchantRequest
	if request.ID = ctx.Param("ID"); request.ID == "" {
		msg := fmt.Sprintf("Failed to parse DeleteMerchant req[%v]", ctx)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := service2.DeleteMerchant(request); err != nil {
		msg := fmt.Sprintf("Failed to handle DeleteMerchant req[%v]: %v", request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
