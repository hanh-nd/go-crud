package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"unique;not null" json:"name"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Permissions []Permission   `gorm:"many2many:role_permissions;foreignKey:ID;joinForeignKey:RoleID;references:ID;joinReferences:PermissionID" json:"permissions"`
}
