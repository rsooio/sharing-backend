package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title     string `gorm:"unique"`
	Abstract  string
	Content   string
	CoverPath string
	AuthorID  uint
	TypeID    uint
}
