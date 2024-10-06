package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ginDemo/config"
	"ginDemo/entity"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.LOG_FILE_PATH
	logFileName := config.LOG_FILE_NAME

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
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
		bodyLogWriter := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = bodyLogWriter

		//开始时间
		startTime := time.Now()
		//处理请求
		ctx.Next()
		//结束时间
		endTime := time.Now()

		//执行时间
		costTime := endTime.Sub(startTime)

		resBody := bodyLogWriter.body.String()

		var resCode string
		var resMsg string
		var resData interface{}

		if resBody != "" {
			res := entity.Result{}
			err := json.Unmarshal([]byte(resBody), &res)
			if err == nil {
				resCode = strconv.Itoa(res.Code)
				resMsg = res.Message
				resData = res.Data
			}
		}

		//请求方式
		// method := ctx.Request.Method

		if ctx.Request.Method == "POST" {
			ctx.Request.ParseForm()
		}
		//请求路由
		// reqUri := ctx.Request.RequestURI
		// //状态码
		// statusCode := ctx.Writer.Status()
		// //请求IP
		// clientIp := ctx.ClientIP()

		//日志格式
		logger.WithFields(logrus.Fields{
			"request_method":       ctx.Request.Method,
			"request_uri":          ctx.Request.RequestURI,
			"request_proto":        ctx.Request.Proto,
			"request_useragent":    ctx.Request.UserAgent(),
			"request_referer":      ctx.Request.Referer(),
			"request_post_data":    ctx.Request.PostForm.Encode(),
			"request_client_ip":    ctx.ClientIP(),
			"response_status_code": ctx.Writer.Status(),
			"response_code":        resCode,
			"response_msg":         resMsg,
			"response_data":        resData,
			"cost_time":            costTime,
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
