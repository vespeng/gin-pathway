package models

type User struct {
	Id       int64  `xorm:"pk autoincr 'id'"`
	Username string `xorm:"varchar(255) not null 'username'"`
	Password string `xorm:"varchar(255) not null 'password'"` // 密码，非空
}

// TableName 方法用于返回表名
func (u User) TableName() string {
	return "user"
}
