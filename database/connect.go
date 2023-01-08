package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"hanhngo.me/m/config"
	"hanhngo.me/m/model"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Get("DB_HOST"),
		config.Get("DB_PORT", "5432"),
		config.Get("DB_USER"),
		config.Get("DB_PASS"),
		config.Get("DB_NAME"),
		config.Get("DB_SSL_MODE", "disable"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db
	fmt.Println("Connected to the database")
	db.AutoMigrate(
		&model.Permission{},
		&model.Role{},
		&model.RoleGroup{},
		&model.User{},
		&model.UserGroup{},
		&model.RolePermission{},
		&model.RoleGroupRole{},
		&model.UserRole{},
		&model.UserRoleGroup{},
		&model.UserGroupUser{},
	)
	setupRelationships(db)
}

func setupRelationships(db *gorm.DB) {
	db.SetupJoinTable(&model.Role{}, "Permissions", &model.RolePermission{})
	db.SetupJoinTable(&model.RoleGroup{}, "Roles", &model.RoleGroupRole{})
	db.SetupJoinTable(&model.User{}, "Roles", &model.UserRole{})
	db.SetupJoinTable(&model.User{}, "RoleGroups", &model.UserRoleGroup{})
	db.SetupJoinTable(&model.User{}, "UserGroups", &model.UserGroupUser{})
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
