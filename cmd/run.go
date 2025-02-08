package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"vesgo/api"
	"vesgo/config"
	"vesgo/initializers"
	"vesgo/middleware"
)

/*
@Author : Vespeng
@Time   : 2025/2/7
@Desc   : 启动服务
*/

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

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	api.SetupRoutes(r, initializers.Engine)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		log.Error("服务启动错误: %v", err)
		return
	}
}

// InitializeAll 初始化所有模块
func InitializeAll() error {
	err := initializers.InitializeDB()
	if err != nil {
		return fmt.Errorf("MySQL初始化错误: %v", err)
	}
	/*err = initializers.InitializeRedis()
	if err != nil {
		return fmt.Errorf("redis初始化错误: %v", err)
	}*/
	err = initializers.InitializeLogger()
	if err != nil {
		return fmt.Errorf("日志初始化错误: %v", err)
	}

	return nil
}
