package services

import (
	"gin-pathway/internal/models"
	"gin-pathway/internal/repositories"
	"github.com/go-xorm/xorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(engine *xorm.Engine) *UserService {
	return &UserService{userRepo: repositories.NewUserRepository(engine)}
}

func (us *UserService) GetUsers() ([]*models.User, error) {
	return us.userRepo.GetUsers()
}
