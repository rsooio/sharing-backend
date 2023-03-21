package model

type UserRole struct {
	UserID uint `gorm:"primarykey"`
	RoleID uint `gorm:"primarykey"`
}
