package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"uniqueIndex;not null" json:"username"`
	Password     string         `gorm:"not null" json:"password"`
	Email        string         `json:"email"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Roles        []Role         `gorm:"many2many:user_roles" json:"roles"`
	RoleGroups   []RoleGroup    `gorm:"many2many:user_role_groups" json:"roleGroups"`
	UserGroups   []*UserGroup   `gorm:"many2many:user_group_users" json:"userGroups"`
	ManageGroups []UserGroup    `gorm:"foreignKey:ManagerID" json:"manageGroups"`
}
