package group

import "gorm.io/gorm"

type UserGroup struct {
	gorm.Model
	Name       string
	AuthRoleID uint
	AuthRole   AuthRole
}

type AuthRole struct {
	gorm.Model
	Roles []Role
}

type Role struct {
	gorm.Model
	Name       string
	AuthRoleID uint
}
