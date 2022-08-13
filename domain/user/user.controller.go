package user

import (
	"gin_gorm/util/get"
	"gin_gorm/util/post"
	"github.com/gin-gonic/gin"
)

func GetName(c *gin.Context) {
	var token, exists = c.Get("token")

	if exists {
		u := getUser(token.(string))
		get.Ok(c, u)
	} else {
		get.Bad(c, "not principal")
	}
}

type saveDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PostSave(c *gin.Context) {
	var dto = saveDto{}
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
