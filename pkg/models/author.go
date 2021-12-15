package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model

	Name     string `gorm:"index;type:varchar(32);comment:用户昵称"`
	Password string `gorm:"type:varchar(32);comment:用户密码"`
}

func init() {
	DB.AddTable(&Author{})
}
