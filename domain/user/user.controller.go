package user

import (
	"gin_gorm/util/get"
	"github.com/gin-gonic/gin"
)

func GetName(c *gin.Context) {
	u := User{}.getUser(c)
	if u == "" {
		get.Bad(c, "not principal")
	} else {
		get.Ok(c, User{}.getUser(c))
	}
}

func Login(c *gin.Context) {
}
