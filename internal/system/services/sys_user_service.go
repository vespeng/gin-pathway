package services

import (
	"github.com/go-xorm/xorm"
	"vesgo/internal/system/models"
)

/*
@Author : Vespeng
@Time   : 2025/2/16
@Desc   : 用户服务
*/

type SysUserService struct {
	engine *xorm.Engine
}

func NewSysUserService(engine *xorm.Engine) *SysUserService {
	return &SysUserService{engine: engine}
}

func (us *SysUserService) GetSysUsers() ([]models.SysUser, error) {
	var users []models.SysUser
	// 通过 UserEntity 的 TableName 方法获取表名
	err := us.engine.Table(models.SysUser{}.TableName()).Find(&users)
	return users, err
}
