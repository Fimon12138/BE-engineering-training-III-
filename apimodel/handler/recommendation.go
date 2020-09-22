package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tickethub_service/service"
	"tickethub_service/util/errors"
	"tickethub_service/util/log"
)

func ListRecommendations(ctx *gin.Context) {
	resp, err := service.ListRecommendation()
	if err != nil {
		msg := fmt.Sprintf("Failed to get recommendationList: %v", err)
		log.Errorf(msg)
		errors.AbortWithWriteErrorResponse(ctx, errors.InternalError(msg))
	}

	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}
