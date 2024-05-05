package routes

import (
	"fmt"
	"gin-demo/controller"

	"github.com/gin-gonic/gin"
)

func initMiddleware(ctx *gin.Context) {
	fmt.Println("路由分组中间件")
	ctx.Next()
}

func AdminRoutesInit(router *gin.Engine) {
	adminRouter := router.Group("/admin", initMiddleware)
	//另外一种写法
	// adminRouter.Use(initMiddleware)
	{
		// adminRouter.GET("user", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "用户")
		// })

		// adminRouter.GET("/news", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "news")
		// })

		adminRouter.GET("/user", controller.UserController{}.Index)
		adminRouter.POST("/user/add", controller.UserController{}.AddUser)
		adminRouter.GET("/news", controller.NewsController{}.Index)
		adminRouter.GET("/user/list", controller.UserController{}.List)
		adminRouter.GET("/user/edit", controller.UserController{}.Edit)
		adminRouter.POST("/user/save", controller.UserController{}.Save)
		adminRouter.GET("/user/delete", controller.UserController{}.Delete)
		adminRouter.POST("/user/search", controller.UserController{}.Search)
	}
}
