package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"column:name"`
	Email     string         `gorm:"column:email"`
	Password  string         `gorm:"column:password"`
	RoleID    uint           `gorm:"column:role_id not null; default:1"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Role      Role           `gorm:"foreignKey:RoleID"`
	Plans     []Plan         `gorm:"foreignKey:UserID"`
}
type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
