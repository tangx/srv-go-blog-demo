package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title    string
	Content  string
	AuthorID int
}

func init() {
	DB.AddTable(&Post{})
}
