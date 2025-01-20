package services

import (
	"github.com/go-xorm/xorm"
	"vesgo/internal/system/models"
)

type UserService struct {
	engine *xorm.Engine
}

func NewUserService(engine *xorm.Engine) *UserService {
	return &UserService{engine: engine}
}

func (us *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	// 通过 UserEntity 的 TableName 方法获取表名
	err := us.engine.Table(models.User{}.TableName()).Find(&users)
	return users, err
}
