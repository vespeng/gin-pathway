package app

import (
	"gin-pathway/config"
	"github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// InitializeLogger 设置日志输出并初始化日志文件
func InitializeLogger() error {
	// 设置日志格式
	switch config.Conf.Log.Format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	default:
		log.SetFormatter(&log.JSONFormatter{})
	}

	// 设置日志级别
	switch config.Conf.Log.Level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	// 设置打印调用信息
	log.SetReportCaller(config.Conf.Log.ReportCaller)

	logDir := "../logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Error("创建日志目录失败: %v", err)
		return err
	}

	// 设置日志输出，按天切割
	logFilePath := logDir + "/app.%Y%m%d.log"
	writer, err := rotatelogs.New(
		logFilePath,
		rotatelogs.WithLinkName(logDir+"/app.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 保留7天
		rotatelogs.WithRotationTime(24*time.Hour), // 每天切割一次
	)
	if err != nil {
		log.Error("设置日志输出失败: %v", err)
		return err
	}

	multiWriter := io.MultiWriter(os.Stdout, writer)
	log.SetOutput(multiWriter)

	return nil
}
