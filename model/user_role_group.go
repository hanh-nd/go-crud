package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRoleGroup struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"primaryKey" json:"userId"`
	RoleGroupID uint           `gorm:"primaryKey" json:"roleGroupId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
