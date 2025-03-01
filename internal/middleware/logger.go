package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 继续处理请求
		c.Next()

		// 记录请求结束时间
		end := time.Now()
		latency := end.Sub(start)

		// 获取请求方法和路径
		method := c.Request.Method
		path := c.Request.URL.Path

		// 获取响应状态码
		statusCode := c.Writer.Status()

		// 记录日志
		logrus.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"method":      method,
			"path":        path,
		}).Info("Request processed")
	}
}
