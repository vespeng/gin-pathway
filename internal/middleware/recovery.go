package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 记录 Panic
				log.Warnf("Recovered from panic: %v", r)

				// 将 Panic 转换为统一的错误格式
				err := NewAppError(
					http.StatusInternalServerError,
					500,
					"Internal server error",
				)

				// 返回统一的错误响应
				c.JSON(err.HTTPStatus, gin.H{
					"code":    err.Code,
					"message": err.Message,
				})

				// 中断请求
				c.Abort()
			}
		}()
		c.Next()
	}
}
