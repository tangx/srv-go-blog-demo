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
	AuthorID  uint   `gorm:"index:cidx_author_post,priority:10"`
	PostID    uint   `gorm:"index:cidx_author_post,priority:5"`
	LikeOrNot string `gorm:"type:enum('like','dislike','nothing'); default:'nothing'"`
}
