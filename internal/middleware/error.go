package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppError struct {
	HTTPStatus int    // HTTP 状态码（如 404）
	Code       int    // 业务错误码（如 40401）
	Message    string // 错误消息（如 "User not found"）
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(httpStatus, code int, message string) *AppError {
	return &AppError{
		HTTPStatus: httpStatus,
		Code:       code,
		Message:    message,
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			var appErr *AppError
			if errors.As(err, &appErr) {
				// 自定义业务错误
				c.JSON(appErr.HTTPStatus, gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				})
			} else {
				// 非自定义错误（如 Go 原生 error）
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Internal server error",
				})
			}
			c.Abort()
		}
	}
}
