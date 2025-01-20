package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"vesgo/api/controllers/system"
	"vesgo/internal/system/services"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine, engine *xorm.Engine) {
	// 定义用户路由组
	user := r.Group("/system")
	{
		// 创建 UserService 实例
		userService := services.NewUserService(engine)
		// 创建 UserController 实例
		userController := system.NewUserController(userService)

		user.GET("/users", userController.GetUsers)
	}
}
