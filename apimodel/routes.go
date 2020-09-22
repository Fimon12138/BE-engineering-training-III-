package apimodel

import (
	"github.com/gin-gonic/gin"
	"tickethub_service/apimodel/handler"
	"tickethub_service/util/enum"
)

func RegisterRoutes(router *gin.Engine) {
	v2 := router.Group(enum.V1)
	{
		v2.GET("/account/:id", handler.GetAccount)
		v2.POST("/account", handler.CreateAccount)
		v2.PUT("/account", handler.UpdateAccount)
		v2.DELETE("/account/:id", handler.DeleteAccount)
		v2.POST("/account/login", handler.LogIn)

		v2.POST("/comment/list", handler.ListComment)
		v2.POST("/comment", handler.CreateComment)
		v2.PUT("/comment", handler.UpdateComment)

		v2.POST("/favorite/list", handler.ListFavorite)
		v2.POST("/favorite", handler.CreateFavorite)
		v2.DELETE("/favorite/:id", handler.DeleteFavorite)

		v2.GET("/merchant/:id", handler.GetMerchant)
		v2.POST("/merchant", handler.CreateMerchant)
		v2.PUT("/merchant", handler.UpdateMerchant)
		v2.DELETE("/merchant/:id", handler.DeleteMerchant)

		v2.POST("/order/list", handler.ListOrder)
		v2.POST("/order", handler.CreateOrder)
		v2.POST("/order/pay", handler.PayForOrder)
		v2.POST("/order/finished", handler.ListFinishedTicket)

		v2.GET("/ticket:id", handler.GetTicket)
		v2.POST("/ticket/list", handler.ListTicket)

		v2.GET("/user/:id", handler.GetUser)
		v2.POST("/user", handler.CreateUser)
		v2.PUT("/user", handler.UpdateUser)
		v2.DELETE("/user/:id", handler.DeleteUser)

		v2.GET("/zjpay/:id", handler.GetZjPay)
		v2.POST("/zjpay/charge", handler.ChargeMoneyRequest)

		v2.GET("/recommendation", handler.ListRecommendations)
	}

}
