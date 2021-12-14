package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model

	Name     string
	Password string
}

func init() {
	DB.AddTable(&Author{})
}
