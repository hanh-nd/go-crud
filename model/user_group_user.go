package model

import (
	"time"

	"gorm.io/gorm"
)

type UserGroupUser struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserGroupID uint           `gorm:"primaryKey" json:"userGroupId"`
	UserID      uint           `gorm:"primaryKey" json:"userId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
