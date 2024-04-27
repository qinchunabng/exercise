package tools

import (
	"crypto/md5"
	"fmt"
	"time"
)

// 时间戳转为时间格式字符串
func UnitToDate(timestamp int) string {
	time := time.Unix(int64(timestamp), 0)
	return time.Format("2006-01-02 15:04:05")
}

// 日志格式字符串转为时间戳
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetDate() string {
	template := "2006-01-02 15:03:04"
	return time.Now().Format(template)
}

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// md5加密
func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}
