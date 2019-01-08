package models

import "encoding/gob"

type User struct {
	ID         int    `form:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;index:id"`
	Name       string `form:"name" gorm:"column:name;index:name"`
	Email      string `form:"email" gorm:"column:email;primary_key;index:email"`
	Mobile     string `form:"mobile" gorm:"column:mobile;index:mobile"`
	Password   string `form:"password" gorm:"column:password"`
	Salt       string `form:"salt" gorm:"column:salt"`
	Lock       bool   `form:"lock" gorm:"column:lock"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
	LoginTime  int64  `gorm:"column:login_time"`
}

func (User) TableName() string {
	return "t_user"
}

func init() {
	migrate(&User{})
	gob.Register(User{})
}
