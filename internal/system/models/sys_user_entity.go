package models

import (
	"time"
)

/*
@Author : Vespeng
@Time   : 2025/2/16
@Desc   : 系统用户表
*/

type SysUser struct {
	Id          int64     `xorm:"pk autoincr 'id'"`
	UserId      int64     `xorm:"not null 'user_id'"`
	Password    string    `xorm:"varchar(50) not null 'password'"`
	DeptId      int64     `xorm:"'dept_id'"`
	UserName    string    `xorm:"varchar(30) 'user_name'"`
	Email       string    `xorm:"varchar(50) 'email'"`
	PhoneNumber int64     `xorm:"'phone_number'"`
	Sex         string    `xorm:"char(1) 'sex'"`
	Avatar      string    `xorm:"varchar(100) 'avatar'"`
	Status      bool      `xorm:"tinyint(1) 'status'"`
	DelFlag     bool      `xorm:"tinyint(1) 'del_flag'"`
	LoginIp     string    `xorm:"varchar(128) 'login_ip'"`
	LoginDate   time.Time `xorm:"datetime 'login_date'"`
	CreateBy    string    `xorm:"varchar(64) 'create_by'"`
	CreateTime  time.Time `xorm:"datetime 'create_time'"`
	UpdateBy    string    `xorm:"varchar(64) 'update_by'"`
	UpdateTime  time.Time `xorm:"datetime 'update_time'"`
	Remark      string    `xorm:"varchar(500) 'remark'"`
}

// TableName 方法用于返回表名
func (u SysUser) TableName() string {
	return "sys_user"
}
