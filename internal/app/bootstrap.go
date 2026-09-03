package app

import (
	"fmt"
	"gin-pathway/internal/app/config"
	"gin-pathway/internal/app/setup"
	"gin-pathway/internal/biz/user"
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

	err = setupDependencies()
	if err != nil {
		log.Errorf("模块初始化错误: %v", err)
		return
	}

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.ErrorHandler())

	v1 := r.Group("/api/v1")
	user.RegisterRoutes(v1, setup.Engine)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		log.Errorf("服务启动错误: %v", err)
		return
	}
}

// setupDependencies 初始化应用依赖
func setupDependencies() error {
	err := setup.Logger()
	if err != nil {
		return fmt.Errorf("日志初始化错误: %v", err)
	}
	err = setup.DB()
	if err != nil {
		return fmt.Errorf("MySQL初始化错误: %v", err)
	}
	err = setup.Redis()
	if err != nil {
		return fmt.Errorf("redis初始化错误: %v", err)
	}
	return nil
}
