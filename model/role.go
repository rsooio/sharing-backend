package model

import (
	"backend/internal/pkg/res"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Describe  string
	Resources res.ResourceFlag
}
