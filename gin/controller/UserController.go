package controller

import (
	"fmt"
	"gin-demo/models"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (c UserController) List(ctx *gin.Context) {
	users := []models.User{}
	models.GetDB().Find(&users)
	ctx.HTML(http.StatusOK, "admin/users.html", gin.H{
		"users": users,
	})
}

func (c UserController) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("username")
	ctx.HTML(http.StatusOK, "admin/adduser.html", gin.H{
		"username": username,
	})
}

func (c UserController) AddUser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	if username != "" {
		//将用户名保存到session中
		session := sessions.Default(ctx)
		session.Set("username", username)
		session.Save()
	}
	//获取上传图片
	file, err := ctx.FormFile("face")
	if err == nil {
		//获取文件后缀名
		ext := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		//判断文件是否为图片
		if _, ok := allowExtMap[ext]; !ok {
			ctx.String(http.StatusOK, "文件类型不合法")
			return
		}
		//创建图片目录保存图片
		day := models.GetDay()
		dir := "./static/upload/" + day
		if err := os.MkdirAll(dir, 0666); err != nil {
			fmt.Println(err)
		}
		//生成文件名
		fileName := strconv.FormatInt(models.GetUnix(), 10)
		saveDir := path.Join(dir, fileName+ext)
		ctx.SaveUploadedFile(file, saveDir)

		json := gin.H{
			"username": username,
			"message":  "保存成功",
			"path":     saveDir,
		}
		//保存用户信息到数据库
		var user models.User
		if err := ctx.ShouldBind(&user); err == nil {
			log.Printf("user:%#v\n", user)
			user.AddTime = time.Now().Unix()
			result := models.GetDB().Create(&user)
			if result.RowsAffected > 0 {
				json["id"] = user.Id
			}
		} else {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, json)
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message":  "文件上传失败",
			"username": username,
			"error":    err.Error(),
		})
	}

}
