package model

import (
	"time"

	"gorm.io/gorm"
)

type RolePermission struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	RoleID       uint           `gorm:"primaryKey" json:"roleId"`
	PermissionID uint           `gorm:"primaryKey" json:"permissionId"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
