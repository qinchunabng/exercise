package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"ginDemo/config"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
)

// 打印
func Print(i interface{}) {
	fmt.Println("---")
	fmt.Println(i)
	fmt.Println("---")
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// MD5方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 生成签名
func CreateSign(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}
	// sort.Strings(key)
	slices.Sort(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		}
	}
	//自定义前面方法
	sign := MD5(MD5(str) + MD5(config.APP_NAME+config.APP_SECRET))
	return sign
}

// 返回JSON
func RetJson(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}
