package apimodel

import (
	"github.com/gin-gonic/gin"
	"tickethub_service/apimodel/handler"
	"tickethub_service/util/enum"
)

func RegisterRoutes(router *gin.Engine) {
	v2 := router.Group(enum.V1)
	{
		v2.POST("/account/info", handler.GetAccount)
		v2.POST("/account", handler.CreateAccount)
		v2.PUT("/account", handler.UpdateAccount)
		v2.POST("/account/delete", handler.DeleteAccount)
		v2.POST("/account/login", handler.LogIn)
		v2.POST("/account/signup", handler.SignUp)

		v2.POST("/comment/list", handler.ListComment)
		v2.POST("/comment", handler.CreateComment)
		v2.PUT("/comment", handler.UpdateComment)

		v2.POST("/favorite/list", handler.ListFavorite)
		v2.POST("/favorite", handler.CreateFavorite)
		v2.POST("/favorite/delete", handler.DeleteFavorite)

		v2.POST("/merchant/info", handler.GetMerchant)
		v2.POST("/merchant", handler.CreateMerchant)
		v2.PUT("/merchant", handler.UpdateMerchant)
		v2.POST("/merchant/delete", handler.DeleteMerchant)

		v2.POST("/order/list", handler.ListOrder)
		v2.POST("/order", handler.CreateOrder)
		v2.PUT("/order", handler.UpdateOrder)
		v2.POST("/order/delete", handler.DeleteOrder)
		v2.POST("/order/pay", handler.PayForOrder)
		v2.POST("/order/finished", handler.ListFinishedTicket)

		v2.POST("/ticket/info", handler.GetTicket)
		v2.POST("/ticket/list", handler.ListTicket)

		v2.POST("/user/info", handler.GetUser)
		v2.POST("/user", handler.CreateUser)
		v2.POST("/user/update", handler.UpdateUser)
		v2.POST("/user/delete", handler.DeleteUser)

		v2.POST("/zjpay/info", handler.GetZjPay)
		v2.POST("/zjpay/charge", handler.ChargeMoneyRequest)

		v2.GET("/recommendation", handler.ListRecommendations)
	}

}
