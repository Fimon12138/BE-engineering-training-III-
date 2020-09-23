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

func CreateOrder(ctx *gin.Context) {
	var req request.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse CreateOrder req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.CreateOrder(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to CreateOrder by req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func ListOrder(ctx *gin.Context) {
	var req request.ListOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse ListOrder req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckListOrder(&req); err != nil {
		log.Errorf("Failed to validate ListOrder req[%v]: %v", req, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	resp, err := service.ListOrder(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle ListOrder req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func DeleteOrder(ctx *gin.Context) {
	var req request.DeleteOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse DeleteOrder req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := service.DeleteOrder(req); err != nil {
		msg := fmt.Sprintf("Failed to handle DeleteOrder req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}

func PayForOrder(ctx *gin.Context) {
	var req request.PayForOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse PayForOrder req[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := service.PayForOrder(req); err != nil {
		msg := fmt.Sprintf("Failed to handle PayForOrder req[%v]:%v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}
}

func ListFinishedTicket(ctx *gin.Context) {
	var req request.ListFinishedTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse ListFinishedTicket req:[%v]: %v", ctx.Request, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckListFinishedTicket(&req); err != nil {
		log.Errorf("Failed to validate ListFinishedTicket req[%v]: %v", req, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
	}

	resp, err := service.ListFinishedTicket(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle ListFinishedTicket req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}
