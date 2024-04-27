package controller

import (
	"github.com/gin-gonic/gin"
)

type NewsController struct {
	BaseController
}

func (c NewsController) Index(ctx *gin.Context) {
	c.Success(ctx)
}
