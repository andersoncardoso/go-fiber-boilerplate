package models

import (
	"fmt"
	"myapp/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect with database
func GetDB() *gorm.DB {
	var err error

	cfg := configs.Load()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	return db
}
