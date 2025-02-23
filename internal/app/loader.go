package app

import (
	"fmt"
	"gin-pathway/config"
	"gin-pathway/internal/api/v1"
	"gin-pathway/internal/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Start 启动服务
func Start() {
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

	r := gin.New()
	r.Use(middleware.Recovery())
	r.Use(middleware.ErrorHandler())
	v1.SetupRoutes(r, Engine)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		log.Error("服务启动错误: %v", err)
		return
	}
}

// InitializeAll 初始化所有模块
func InitializeAll() error {
	err := InitializeLogger()
	if err != nil {
		return fmt.Errorf("日志初始化错误: %v", err)
	}
	err = InitializeDB()
	if err != nil {
		return fmt.Errorf("MySQL初始化错误: %v", err)
	}
	/*err = initializers.InitializeRedis()
	if err != nil {
		return fmt.Errorf("redis初始化错误: %v", err)
	}*/

	return nil
}
