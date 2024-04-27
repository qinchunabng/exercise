package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutesInit(router *gin.Engine) {
	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"username": "张三",
				"age":      20,
			})
		})
		apiRoute.GET("/news", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"title": "这是新闻",
			})
		})
	}
}
