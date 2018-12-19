package models

import "time"

type TestModule struct {
	ID        int        `form:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;index:id"`
	Name      string     `form:"name" gorm:"column:name;index:name"`
	CreatedAt time.Time  `gorm:"column:create_at"`
	UpdatedAt time.Time  `gorm:"column:update_at"`
	DeletedAt *time.Time `gorm:"column:delete_at"`
}

func (TestModule) TableName() string {
	return "t_test_module"
}

func init() {
	migrate(&TestModule{})
}
