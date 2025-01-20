package initializers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"vesgo/api"
	"vesgo/config"
)

// InitializeAll 初始化所有模块
func InitializeAll() error {
	err := InitializeMySQL()
	if err != nil {
		return fmt.Errorf("MySQL初始化错误: %v", err)
	}
	/*err = InitializeRedis()
	if err != nil {
		return fmt.Errorf("redis初始化错误: %v", err)
	}*/
	err = InitializeLogger()
	if err != nil {
		return fmt.Errorf("日志初始化错误: %v", err)
	}

	return nil
}

// Run 启动服务
func Run() {
	err := config.LoadConfig()
	if err != nil {
		log.Error("配置文件加载错误: %v", err)
		return
	}

	err = InitializeAll()
	if err != nil {
		log.Error("模块初始化错误: %v", err)
		return
	}

	r := gin.Default()
	api.SetupRoutes(r, Engine)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		log.Error("服务启动错误: %v", err)
		return
	}
}
