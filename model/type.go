package model

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Describe string
	ParentID uint
}
