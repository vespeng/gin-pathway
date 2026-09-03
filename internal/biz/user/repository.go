package user

import "github.com/go-xorm/xorm"

type UserRepository struct {
	engine *xorm.Engine
}

func NewUserRepository(engine *xorm.Engine) *UserRepository {
	return &UserRepository{engine: engine}
}

func (r *UserRepository) GetUsers() ([]*User, error) {
	var users []*User
	err := r.engine.Table(User{}.TableName()).Find(&users)
	return users, err
}
