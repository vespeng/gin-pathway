package repositories

import (
	models "gin-pathway/internal/model"
	"github.com/go-xorm/xorm"
)

type UserRepository struct {
	engine *xorm.Engine
}

func NewUserRepository(engine *xorm.Engine) *UserRepository {
	return &UserRepository{engine: engine}
}

// GetUsers 获取所有用户
func (r *UserRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.engine.Table(models.User{}.TableName()).Find(&users)
	return users, err
}
