package user

import (
	"gin_gorm/util/get"
	"gin_gorm/util/post"
	"gin_gorm/util/security"
	"github.com/gin-gonic/gin"
)

func GetName(c *gin.Context) {
	var token = c.GetHeader("Cookie")

	var strToken = token[len("token="):]

	if strToken == "" {
		get.Bad(c, "not principal")
	} else {
		u := getUser(strToken)
		get.Ok(c, u)
	}

}

type namePwDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PostSave(c *gin.Context) {
	var dto = namePwDto{}
	var err1 = c.BindJSON(&dto)

	var name = dto.Name
	var pw = dto.Password

	if err1 == nil {
		id, err1 := save(User{Name: name, Password: pw})
		if err1 != nil {
			post.Bad(c, "fail save err : "+err1.Error())
		} else {
			post.Ok(c, id)
		}
	} else {
		post.Bad(c, "empty name, pw")
	}
}

func PostLogin(c *gin.Context) {
	var dto = namePwDto{}
	var err1 = c.BindJSON(&dto)

	var name = dto.Name
	var pw = dto.Password

	if err1 == nil {
		token := login(User{Name: name, Password: pw})

		if token == "" {
			get.Bad(c, "")
		} else {
			security.Login(c, token)
		}
	} else {
		post.Bad(c, "err?"+err1.Error())
	}

}
