package user

import (
	"gin_gorm/util/auth"
	"github.com/cristalhq/jwt/v4"
	"github.com/gin-gonic/gin"
)

func (u User) getUser(ctx *gin.Context) string {
	name, exists := ctx.Get("User")
	if exists {
		return name.(User).Name
	}
	return ""
}

func getId() string {
	rawToken := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJnby1naW4tcHJvamVjdC1jb3JlIiwiYXVkIjoiMTEiLCJpc19hZG1pbiI6ZmFsc2UsImVtYWlsIjoidGVzdEBhc2tlLmNvLmtyIn0.PTXUo32VfBNI0QX71vHyVklXpeO0uAvK_dEKtsxGTnk")

	key := []byte(`jwt_!test`)
	verifier, _ := jwt.NewVerifierHS(jwt.HS256, key)

	token, _ := jwt.Parse(rawToken, verifier)

	return auth.ParseToken(token)
}

func (u User) Login() string {
	str, _ := auth.CreateToken(u.Id)
	return str
}
