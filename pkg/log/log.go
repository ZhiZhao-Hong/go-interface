package log

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

func Get() *zap.SugaredLogger {
	return logger
}

type Conf struct {
	Path       string        `json:"path"`
	Level      zapcore.Level `json:"level"`
	MaxSize    int           `json:"maxSize"`
	MaxBackups int           `json:"maxBackups"`
	MaxAge     int           `json:"maxAge"`
	Compress   bool          `json:"compress"`
}

func (c *Conf) InitLogger() {
	// 打印指定级别的日志
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if c.Level.String() == "" {
			return true
		}

		return lvl >= c.Level
	})

	hook := lumberjack.Logger{
		Filename:   c.Path,       //文件名 路径+名称 ./logs/a.log
		MaxSize:    c.MaxSize,    // mb
		MaxBackups: c.MaxBackups, // 备份
		MaxAge:     c.MaxAge,     // 保存天数
		Compress:   c.Compress,   // 是否压缩
	}

	fileWriter := zapcore.AddSync(&hook)
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		// 打印在控制台
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),

		// 打印在文件中
		zapcore.NewCore(consoleEncoder, fileWriter, lowPriority),
	)

	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = log.Sugar()
}

func GinToLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func StringToLogLevel(level string) zapcore.Level {
	switch level {
	case "fatal":
		return zap.FatalLevel
	case "error":
		return zap.ErrorLevel
	case "warn":
		return zap.WarnLevel
	case "warning":
		return zap.WarnLevel
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	}
	return zap.DebugLevel
}
