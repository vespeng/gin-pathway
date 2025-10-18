package v1

import (
	"gin-pathway/internal/controller"
	"gin-pathway/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine, engine *xorm.Engine) {
	// 定义用户路由组
	user := r.Group("/user")
	{
		// 创建 UserService 实例
		UserService := service.NewUserService(engine)
		// 创建 UserController 实例
		UserController := controller.NewUserController(UserService)

		user.GET("/", UserController.GetUsers)
	}
}
