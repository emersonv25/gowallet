package router

import (
	docs "github.com/emersonv25/gowallet/docs"
	"github.com/emersonv25/gowallet/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/wallet", func(ctx *gin.Context) { handler.GetWallet(ctx) })
		v1.POST("/wallet", func(ctx *gin.Context) { handler.CreateWalletHandler(ctx) })
		v1.PUT("/wallet", func(ctx *gin.Context) { handler.UpdateWalletHandler(ctx) })
		v1.DELETE("/wallet", func(ctx *gin.Context) { handler.DeleteWalletHandler(ctx) })

		v1.GET("/transaction", func(ctx *gin.Context) { handler.GetTransactions(ctx) })
		v1.POST("/transaction", func(ctx *gin.Context) { handler.CreateTransaction(ctx) })
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
