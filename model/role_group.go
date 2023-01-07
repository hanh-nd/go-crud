package model

import (
	"time"

	"gorm.io/gorm"
)

type RoleGroup struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Roles     []Role         `gorm:"many2many:role_group_roles" json:"roles"`
}
