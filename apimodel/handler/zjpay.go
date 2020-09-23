package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tickethub_service/apimodel/request"
	"tickethub_service/service"
	"tickethub_service/util/errors"
	"tickethub_service/util/log"
)

func GetZjPay(ctx *gin.Context) {
	var req request.GetZjpayRequest

	if req.ID = ctx.Param("id"); req.ID == "" {
		msg := fmt.Sprintf("Failed to parse GetZjpay req:[%v]", &ctx)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.GetZjpay(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle GetZjpay by req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func ChargeMoneyRequest(ctx *gin.Context) {
	var req request.ChargeMoneyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse ChargeMoneyRequest ctx[%v]", &ctx)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	err := service.ChargeMoney(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle ChargeMoneyRequest[%v]: [%v]", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
