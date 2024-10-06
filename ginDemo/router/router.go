package router

import (
	common "ginDemo/common/function"
	v1 "ginDemo/controller/v1"
	v2 "ginDemo/controller/v2"
	"ginDemo/middleware/logger"
	"ginDemo/middleware/recover"
	"ginDemo/validator/member"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitRouter(r *gin.Engine) {
	r.Use(logger.LoggerToFile(), recover.Recover())
	r.GET("/sn", SignDemo)

	//v1版本
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}

	//v2版本
	GroupV2 := r.Group("/v2")
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}

	//绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", member.NameValid)
	}
}

func SignDemo(c *gin.Context) {
	ts := common.GetTimeUnix()
	res := map[string]interface{}{}
	params := url.Values{
		"name":  []string{"a"},
		"price": []string{"10"},
		"ts":    []string{ts},
	}
	res["sn"] = common.CreateSign(params)
	res["ts"] = ts
	common.RetJson("200", "", res, c)
}
