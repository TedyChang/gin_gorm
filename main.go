package main

import (
	"gin_gorm/api"
	"gin_gorm/domain/user"
	db "gin_gorm/gorm"
	"log"
)

func main() {
	err1 := db.Run()

	if err1 != nil {
		log.Panicf("err :\n%v", err1)
	}

	err2 := db.DB.AutoMigrate(&user.User{})
	if err2 != nil {
		log.Panicf("migration error\nerr :\n%v", err2)
	}

	api.RestApi()
}
