package router

import (
	"github.com/emersonv25/gowallet/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/client", func(ctx *gin.Context) { handler.GetClientHandler(ctx) })
		v1.POST("/client", func(ctx *gin.Context) { handler.CreateClientHandler(ctx) })
		v1.PUT("/client", func(ctx *gin.Context) { handler.UpdateClientHandler(ctx) })
		v1.DELETE("/client", func(ctx *gin.Context) { handler.DeleteClientHandler(ctx) })
	}
}
