package app

import (
	"fmt"
	"gin-pathway/internal/api/v1"
	"gin-pathway/internal/app/config"
	"gin-pathway/internal/app/initializer"
	"gin-pathway/internal/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Start 启动服务
func Start() {
	err := config.LoadConfig()
	if err != nil {
		log.Errorf("配置文件加载错误: %v", err)
		return
	}

	err = InitializeAll()
	if err != nil {
		log.Errorf("模块初始化错误: %v", err)
		return
	}

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.ErrorHandler())
	v1.SetupRoutes(r, initializer.Engine)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		log.Errorf("服务启动错误: %v", err)
		return
	}
}

// InitializeAll 初始化所有模块
func InitializeAll() error {
	err := initializer.InitializeLogger()
	if err != nil {
		return fmt.Errorf("日志初始化错误: %v", err)
	}
	err = initializer.InitializeDB()
	if err != nil {
		return fmt.Errorf("MySQL初始化错误: %v", err)
	}
	err = initializer.InitializeRedis()
	if err != nil {
		return fmt.Errorf("redis初始化错误: %v", err)
	}
	return nil
}
