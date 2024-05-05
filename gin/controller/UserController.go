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

func (c UserController) Edit(ctx *gin.Context) {
	idStr, ok := ctx.GetQuery("id")
	user := models.User{}
	if ok {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.String(http.StatusBadRequest, "参数错误")
			return
		}
		models.GetDB().Where("id=?", id).First(&user)
	}

	ctx.HTML(http.StatusOK, "admin/edit_user.html", &user)
}

func (c UserController) Save(ctx *gin.Context) {
	tx := models.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("err:%v", r)
			tx.Rollback()
			ctx.String(http.StatusAccepted, "更新失败")
		}
	}()
	if err := tx.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		tx.Rollback()
		ctx.String(http.StatusBadRequest, "参数错误")
		return
	}

	result := tx.Save(&user)
	// n := 0
	// i := 10 / n
	// log.Println(i)
	if result.RowsAffected > 0 {
		tx.Commit()
		ctx.String(http.StatusOK, "更新成功")
	} else {
		tx.Rollback()
		ctx.String(http.StatusAccepted, "更新失败")
	}
}

func (c UserController) Delete(ctx *gin.Context) {
	idStr, ok := ctx.GetQuery("id")
	if !ok {
		ctx.String(http.StatusBadRequest, "参数有误")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "参数有误")
		return
	}
	result := models.GetDB().Delete(&models.User{}, id)
	if result.RowsAffected > 0 {
		ctx.String(http.StatusOK, "删除成功")
	} else {
		ctx.String(http.StatusAccepted, "删除失败")
	}
}

func (c UserController) Search(ctx *gin.Context) {
	var userArg models.UserArgument
	if err := ctx.ShouldBind(&userArg); err != nil {
		ctx.String(http.StatusBadRequest, "参数有误")
		return
	}
	log.Printf("%s\n", fmt.Sprintf("%v%v%v", "%", userArg.Username, "%"))
	var users []models.User
	models.GetDB().Raw("select * from user where username like ?", "%"+userArg.Username+"%").Find(&users)
	log.Println(users)
	ctx.HTML(http.StatusOK, "admin/users.html", gin.H{
		"users": users,
	})
}
