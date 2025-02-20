package v1

import (
	"gin-pathway/internal/controllers"
	"gin-pathway/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine, engine *xorm.Engine) {
	// 定义用户路由组
	user := r.Group("/user")
	{
		// 创建 UserService 实例
		UserService := services.NewUserService(engine)
		// 创建 UserController 实例
		UserController := controllers.NewUserController(UserService)

		user.GET("/", UserController.GetUsers)
	}
}
