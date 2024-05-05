package models

import (
	"crypto/md5"
	"fmt"
	"log"
	"time"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func NewDB(cfg *ini.File) {
	mysqlSection := cfg.Section("mysql")
	host := mysqlSection.Key("ip").String()
	port := mysqlSection.Key("port").String()
	database := mysqlSection.Key("database").String()
	username := mysqlSection.Key("user").String()
	passwd := mysqlSection.Key("password").String()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, passwd, host, port, database)
	log.Printf("dsn:%v\n", dsn)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

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

func GetDB() *gorm.DB {
	if db == nil {
		panic(fmt.Errorf("db has not been initialized"))
	}
	return db
}
