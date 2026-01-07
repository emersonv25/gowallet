package router

import (
	"github.com/emersonv25/gowallet/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/wallet", func(ctx *gin.Context) { handler.GetWallet(ctx) })
		v1.POST("/wallet", func(ctx *gin.Context) { handler.CreateWalletHandler(ctx) })
		v1.PUT("/wallet", func(ctx *gin.Context) { handler.UpdateWalletHandler(ctx) })
		v1.DELETE("/wallet", func(ctx *gin.Context) { handler.DeleteWalletHandler(ctx) })

		v1.GET("/transaction", func(ctx *gin.Context) { handler.GetTransactions(ctx) })
		v1.POST("/transaction", func(ctx *gin.Context) { handler.CreateTransaction(ctx) })
	}
}
