package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

// RegisterRoutes 注册用户模块的路由
func RegisterRoutes(rg *gin.RouterGroup, engine *xorm.Engine) {
	repo := NewUserRepository(engine)
	svc := NewUserService(repo)
	handler := NewUserHandler(svc)

	userGroup := rg.Group("/user")
	{
		userGroup.GET("/", handler.GetUsers)
	}
}
