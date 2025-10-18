package service

import (
	"gin-pathway/internal/model"
	"gin-pathway/internal/repository"

	"github.com/go-xorm/xorm"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(engine *xorm.Engine) *UserService {
	return &UserService{userRepo: repository.NewUserRepository(engine)}
}

func (us *UserService) GetUsers() ([]*model.User, error) {
	return us.userRepo.GetUsers()
}
