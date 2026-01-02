package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "pong",
			})
		})
	}
}
