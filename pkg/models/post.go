package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title        string `gorm:"type:varchar(128);index"`
	Content      string `gorm:"longtext"`
	ReadingCount int    `gorm:"uint"`
	AuthorID     int    `gorm:"bigint"`
	Author       Author `gorm:"foreignKey:AuthorID"`
}

func init() {
	DB.AddTable(&Post{})
}
