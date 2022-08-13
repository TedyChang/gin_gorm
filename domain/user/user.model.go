package user

import (
	db "gin_gorm/gorm"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func (u *User) FindById() {
	db.DB.First(&u, u.ID)
}

func (u *User) Create() {
	db.DB.Create(&u)
}
