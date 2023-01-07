package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"primaryKey" json:"userId"`
	RoleID    uint           `gorm:"primaryKey" json:"roleId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
