package middleware

import (
	"fmt"
	"ginDemo/config"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.LOG_FILE_PATH
	logFileName := config.LOG_FILE_NAME

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置rotatelogs
	logWriter, err := rotatelogs.New(
		//分割后的文件名称
		fileName+".%Y%m%d.log",
		//生成软连接，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		//设置最大保存时间（7天）
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//设置日志切割时间间隔（1天）
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//新增钩子
	logger.AddHook(lfHook)

	return func(ctx *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		ctx.Next()
		//结束时间
		endTime := time.Now()

		//执行时间
		costTime := endTime.Sub(startTime)

		//请求方式
		method := ctx.Request.Method
		//请求路由
		reqUri := ctx.Request.RequestURI
		//状态码
		statusCode := ctx.Writer.Status()
		//请求IP
		clientIp := ctx.ClientIP()

		//日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": costTime,
			"client_ip":    clientIp,
			"req_method":   method,
			"request_uri":  reqUri,
		}).Info()
	}
}

// 记录日志到MongoDB
func LoggerToMongoDB() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// 记录日志到ES
func LoggerToES() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// 记录日志到MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
