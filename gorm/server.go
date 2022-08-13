package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Run() error {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=9920 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db
	return err
}
