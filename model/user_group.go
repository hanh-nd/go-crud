package model

import (
	"time"

	"gorm.io/gorm"
)

type UserGroup struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"unique;not null" json:"name"`
	ManagerID uint           `json:"managerId"`
	Manager   User           `gorm:"foreignKey:ManagerID" json:"manager"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Users     []*User        `gorm:"many2many:user_group_users" json:"users"`
}
