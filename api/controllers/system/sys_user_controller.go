package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vesgo/internal/system/services"
)

/*
@Author : Vespeng
@Time   : 2025/2/16
@Desc   : 系统用户控制器
*/

type SysUserController struct {
	sysUserService *services.SysUserService
}

func NewSysUserController(sysUserService *services.SysUserService) *SysUserController {
	return &SysUserController{sysUserService: sysUserService}
}

func (uc *SysUserController) GetSysUsers(c *gin.Context) {
	users, err := uc.sysUserService.GetSysUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
