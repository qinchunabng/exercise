package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMember(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"v2": "AddMember",
	})
}
