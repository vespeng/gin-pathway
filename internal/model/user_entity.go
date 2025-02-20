package models

type User struct {
	Id          int64  `xorm:"pk autoincr 'id'"`
	UserId      int64  `xorm:"not null 'user_id'"`
	Password    string `xorm:"varchar(50) not null 'password'"`
	UserName    string `xorm:"varchar(30) 'user_name'"`
	Email       string `xorm:"varchar(50) 'email'"`
	PhoneNumber int64  `xorm:"'phone_number'"`
	Sex         string `xorm:"char(1) 'sex'"`
	Remark      string `xorm:"varchar(500) 'remark'"`
}

// TableName 方法用于返回表名
func (u User) TableName() string {
	return "user"
}
