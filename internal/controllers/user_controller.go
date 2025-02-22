package controllers

import (
	"gin-pathway/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(UserService *services.UserService) *UserController {
	return &UserController{UserService: UserService}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
