package utils

import (
	"gin-pathway/internal/model"

	"github.com/go-xorm/xorm"
)

// SyncDatabase 同步数据库表结构
func SyncDatabase(engine *xorm.Engine) error {
	return engine.Sync2(new(model.User))
}
