package main

import (
	"errors"
	"limiter-breaker/breaker"
	"limiter-breaker/limiter"
	"limiter-breaker/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Use(middleware.Limiter(limiter.NewLimiter(500*time.Millisecond, 4)))
	r.GET("/ping", middleware.Limiter(limiter.NewLimiter(500*time.Millisecond, 4)), func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	b := breaker.NewBreaker(4, 4, 2, time.Second*15)
	r.GET("/ping1", func(ctx *gin.Context) {
		err := b.Exec(func() error {
			value, _ := ctx.GetQuery("value")
			if value == "a" {
				return errors.New("value为a返回错误")
			}
			return nil
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong1",
		})
	})

	r.Run(":8080")
}
