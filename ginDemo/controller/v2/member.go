package v2

import (
	"ginDemo/entity"
	"ginDemo/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMember(c *gin.Context) {
	res := entity.Result{}
	mem := entity.Member{}
	if err := c.ShouldBind(&mem); err != nil {
		res.SetCode(entity.CODE_FAIL)
		res.SetMessage(err.Error())
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}

	data := map[string]interface{}{
		"name": mem.Name,
		"Age":  mem.Age,
	}
	// res.SetCode(entity.CODE_SUCCESS)
	// res.SetData(data)
	// c.JSON(http.StatusOK, res)

	//改造后
	c.JSON(http.StatusOK, errno.OK.WithData(data))
}
