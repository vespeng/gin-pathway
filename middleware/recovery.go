package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

/*
@Author : Vespeng
@Time   : 2025/2/9
@Desc   : 自定义恢复中间件
*/

func CustomRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 记录 Panic
				log.Printf("Recovered from panic: %v", r)

				// 将 Panic 转换为统一的错误格式
				err := NewAppError(
					http.StatusInternalServerError,
					500,
					"Internal server error",
				)

				// 返回统一的错误响应
				ctx.JSON(err.HTTPStatus, gin.H{
					"code":    err.Code,
					"message": err.Message,
				})

				// 中断请求
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
