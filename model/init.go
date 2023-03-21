package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Article{},
		&Comment{},
		&Type{},
		&User{},
	)

	return db
}
