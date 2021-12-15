package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	PostInfo PostInfo `gorm:"embedded"`
	AuthorID uint     `gorm:"bigint"`
	Author   Author   `gorm:"foreignKey:AuthorID"`
}

type PostInfo struct {
	Title        string `gorm:"type:varchar(128);index"`
	Content      string `gorm:"longtext"`
	ReadingCount int    `gorm:"uint"`
}

func init() {
	DB.AddTable(&Post{})
}
