package main

import (
	"gin_gorm/api"
	"gin_gorm/domain/user"
	"gin_gorm/domain/user/group"
	db "gin_gorm/gorm"
	"log"
)

func main() {
	err1 := db.Run()

	if err1 != nil {
		log.Panicf("err :\n%v", err1)
	}

	AutoMigrate(&user.User{})

	DropTable(&group.AuthRole{})
	AutoMigrate(&group.AuthRole{})

	DropTable(&group.Role{})
	AutoMigrate(&group.Role{})

	DropTable(&group.UserGroup{})
	AutoMigrate(&group.UserGroup{})

	api.RestApi()
}

func AutoMigrate(m interface{}) {
	err1 := db.DB.AutoMigrate(m)
	if err1 != nil {
		log.Panicf("migration error\nerr :\n%v", err1)
	}
}
func DropTable(m interface{}) {
	err1 := db.DB.Migrator().DropTable(m)
	if err1 != nil {
		log.Panicf("migration error\nerr :\n%v", err1)
	}
}
