package models

import "gorm.io/gorm"

func init() {
	DB.AddTable(&Vote{})
}

type Vote struct {
	gorm.Model

	AuthorID  int
	PostID    int
	LikeOrNot string `gorm:"type:enum('like','dislike'); default:'like'"`
	Post      Post   `gorm:"foreignKey:PostID"`
	Author    Author `gorm:"foreeignKey:AuthorID"`
}
