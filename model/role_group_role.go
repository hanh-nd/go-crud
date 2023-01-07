package model

import (
	"time"

	"gorm.io/gorm"
)

type RoleGroupRole struct {
	ID          uint           `gorm:"primaryKey; autoIncrement" json:"id"`
	RoleGroupID uint           `gorm:"primaryKey" json:"roleGroupId"`
	RoleID      uint           `gorm:"primaryKey" json:"roleId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
