package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"vesgo/api/controllers/system"
	"vesgo/internal/system/services"
)

/*
@Author : Vespeng
@Time   : 2025/1/21
@Desc   : 路由设置
*/

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine, engine *xorm.Engine) {
	// 定义用户路由组
	user := r.Group("/system")
	{
		// 创建 UserService 实例
		sysUserService := services.NewSysUserService(engine)
		// 创建 UserController 实例
		sysUserController := system.NewSysUserController(sysUserService)

		user.GET("/users", sysUserController.GetSysUsers)
	}
}
