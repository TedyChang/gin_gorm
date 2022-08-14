package group

import (
	"errors"
	db "gin_gorm/gorm"
	"gorm.io/gorm"
	"log"
)

type UserGroup struct {
	gorm.Model
	GroupName  string
	AuthRoleID uint
	AuthRole   AuthRole
}

type AuthRole struct {
	gorm.Model
	AuthName string
	Roles    []Role
}

type Role struct {
	gorm.Model
	RoleName   string
	AuthRoleID uint
}

func (r Role) Add() (uint, error) {
	if r.RoleName == "" {
		return 0, errors.New("empty role_name")
	}
	db.DB.Create(&r)
	return r.ID, nil
}

func (a AuthRole) AddAuthRole() (uint, error) {
	if a.AuthName == "" {
		return 0, errors.New("empty auth_name")
	}
	db.DB.Create(&a)
	return a.ID, nil
}

func (r Role) AddRole() bool {
	if r.RoleName == "" {
		log.Println("empty role name")
		return false
	}

	var existAuth bool
	err1 := AuthRole{}.Exist(&existAuth)

	if err1 != nil {
		log.Panicf("%v\n", err1)
		return false
	}
	if !existAuth {
		log.Println("not exist role-auth")
		return false
	}

	var exist bool
	err2 := r.FindByAuthAndName(&exist)

	if err2 != nil {
		log.Printf("%v\n", err2)
		return false
	}
	if exist {
		log.Panicf("alread exist \n")
		return false
	}

	db.DB.Create(&r)
	return true
}

func (a AuthRole) Exist(exist *bool) error {
	return db.DB.Model(a).
		Select("count(*) > 0").
		Where("id = ?", a.ID).
		Find(exist).Error
}

func (r Role) FindByAuthAndName(exist *bool) error {
	return db.DB.Model(r).
		Select("count(*) > 0").
		Where("auth_role_id = ? and role_name = ?", r.AuthRoleID, r.RoleName).
		Find(exist).Error
}
