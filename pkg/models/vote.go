package models

import "gorm.io/gorm"

func init() {
	DB.AddTable(&Vote{})
}

type Vote struct {
	gorm.Model

	VoteInfo VoteInfo `gorm:"embedded"`
	Post     Post     `gorm:"foreignKey:PostID"`
	Author   Author   `gorm:"foreeignKey:AuthorID"`
}

type VoteInfo struct {
	AuthorID  uint
	PostID    uint
	LikeOrNot string `gorm:"type:enum('like','dislike','nothing'); default:'nothing'"`
}
