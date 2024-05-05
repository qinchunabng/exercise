package main

import (
	"encoding/xml"
	"fmt"
	"gin-demo/models"
	"gin-demo/routes"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

var (
	cfg *ini.File
)

func init() {
	var err error
	cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("fail to read file: %v\n", err)
		os.Exit(1)
	}
}

type UserInfo struct {
	Name   string   `json:"name"`
	Gender string   `json:"gender"`
	Age    int      `json:"age"`
	Score  float32  `json:"score"`
	Hobby  []string `json:"hobby"`
}

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Article struct {
	Title   string `xml:"title" json:"title"`
	Content string `xml:"content" json:"content"`
}

func formatAsDate(s int64) string {
	t := time.Unix(s, 0)
	return t.Format("2006-01-02 15:04:05")
}

func initMiddleware(ctx *gin.Context) {
	fmt.Println("我是一个中间件")
	start := time.Now().UnixNano()
	//继续执行后面的逻辑
	ctx.Next()
	end := time.Now().UnixNano()
	fmt.Printf("耗时%dns\n", end-start)
}

func main() {
	r := gin.Default()
	//初始化数据集
	models.NewDB(cfg)
	//初始化基于redis的存储引擎
	//参数说明
	//	第1个参数：redis最大的空闲连接数
	//	第2个参数：通信协议
	//	第3个参数：redis地址：ip:port
	//	第4个参数：redis密码
	//	第5个参数：session加密密钥
	redisHost := cfg.Section("redis").Key("ip").String()
	redisPort := cfg.Section("redis").Key("port").String()
	store, _ := redis.NewStore(10, "tcp", fmt.Sprintf("%v:%v", redisHost, redisPort), "", []byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   false,
	})
	r.Use(sessions.Sessions("mysession", store))
	//全局中级间
	r.Use(initMiddleware)
	//前面的static表示路由，后面./static表示路径
	r.Static("/static", "./static")
	//注册全局模板函数
	//必须在加载模板上面
	r.SetFuncMap(template.FuncMap{
		"formatDate": formatAsDate,
	})
	r.LoadHTMLGlob("templates/**/*")

	user := UserInfo{
		Name:   "张三",
		Gender: "男",
		Age:    18,
		Score:  99,
		Hobby: []string{
			"吃饭",
			"睡觉",
			"写代码",
		},
	}
	r.GET("/", initMiddleware, func(ctx *gin.Context) {
		cCp := ctx.Copy()
		//注意：在中间件中启动goroutine不能使用原始上下文
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("Done! in path ", cCp.Request.URL.Path)
		}()
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "前台首页",
			"user":  user,
			"now":   time.Now().Unix(),
		})
	})

	r.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})

	r.GET("/app", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": "hello",
		})
	})

	r.GET("/addUser", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/add_user.html", gin.H{})
	})

	//form表单传参
	r.POST("doAddUser", func(ctx *gin.Context) {
		// username := ctx.PostForm("username")
		// password := ctx.PostForm("password")
		// age := ctx.DefaultPostForm("age", "20")

		// ctx.JSON(http.StatusOK, gin.H{
		// 	"username": username,
		// 	"password": password,
		// 	"age":      age,
		// })

		var user User
		if err := ctx.ShouldBind(&user); err == nil {
			ctx.JSON(http.StatusOK, user)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	//接收xml参数
	r.POST("/xml", func(ctx *gin.Context) {
		b, _ := ctx.GetRawData()
		var article Article
		if err := xml.Unmarshal(b, &article); err == nil {
			ctx.JSON(http.StatusOK, &article)
		} else {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
	})

	routes.AdminRoutesInit(r)
	routes.ApiRoutesInit(r)

	r.Run(":8080")
}
