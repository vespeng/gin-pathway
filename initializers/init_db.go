package initializers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"vesgo/config"
)

var Engine *xorm.Engine

// InitializeMySQL 数据库初始化
func InitializeMySQL() error {
	var err error
	// 创建数据库引擎
	Engine, err = xorm.NewEngine(config.Conf.Database.Driver, config.Conf.Database.Source)
	if err != nil {
		log.Error("数据库初始化失败: %v", err)
		return err
	}

	// 测试数据库连接
	if err = Engine.Ping(); err != nil {
		log.Error("数据库连接失败: %v", err)
		return err
	}

	return nil
}
