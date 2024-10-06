package main

import (
	"fmt"
	"ginDemo/config"
	"ginDemo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	router.InitRouter(engine) //设置路由
	err := engine.Run(config.PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
}
