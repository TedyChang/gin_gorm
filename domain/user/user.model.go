package user

import (
	"gin_gorm/domain/user/group"
	db "gin_gorm/gorm"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Password    string
	UserGroupID uint
	UserGroup   group.UserGroup
}

func (u *User) FindById() {
	db.DB.First(&u, u.ID)
}

func (u *User) Create() {
	db.DB.Create(&u)
}

func (u *User) FindByNamePw() {
	db.DB.Where("name = ? AND password = ?", u.Name, u.Password).First(&u)
}

func (u User) ExistByName(exists *bool) error {
	return db.DB.Model(User{}).
		Select("count(*) > 0").
		Where("name = ?", u.Name).
		Find(exists).
		Error
}
