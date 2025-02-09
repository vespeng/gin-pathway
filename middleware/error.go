package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
@Author : Vespeng
@Time   : 2025/2/9
@Desc   : 错误处理中间件
*/

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
	return func(ctx *gin.Context) {
		ctx.Next()
		// 检查是否有错误
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last()
			var appErr *AppError
			if errors.As(err, &appErr) {
				// 自定义业务错误
				ctx.JSON(appErr.HTTPStatus, gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				})
			} else {
				// 非自定义错误（如 Go 原生 error）
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Internal server error",
				})
			}
			ctx.Abort()
		}
	}
}
