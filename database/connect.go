package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hanhngo.me/m/config"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db
	fmt.Println("Connected to the database")
	db.AutoMigrate()
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
