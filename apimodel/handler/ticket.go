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

func GetTicket(ctx *gin.Context) {
	var req request.GetTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse GetTicket req:[%v]", ctx)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	resp, err := service.GetTicket(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle GetTicket by req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func ListTicket(ctx *gin.Context) {
	var req request.ListTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := fmt.Sprintf("Failed to parse ListTicket req:[%v]", ctx)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	if err := validate.CheckListTicket(&req); err != nil {
		log.Errorf("Failed to validate ListTicket req[%v]: %v", req, err)
		errors.AbortWithWriteErrorResponse(ctx, err)
		return
	}

	resp, err := service.ListTicket(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to handle ListTicket req[%v]: %v", req, err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}
