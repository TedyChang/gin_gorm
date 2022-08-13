package auth

import (
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v4"
	"log"
)

type userClaims struct {
	jwt.RegisteredClaims
	IsAdministrator bool   `json:"is_admin"`
	Email           string `json:"email"`
}

func CreateToken(id uint) (string, error) {
	key := []byte("jwt_!test")
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	checkErr(err)
	builder := jwt.NewBuilder(signer)

	claims := userClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience: []string{fmt.Sprintf("%d", id)},
			ID:       "go-gin-project-core",
		},
		IsAdministrator: false,
		Email:           "test@aske.co.kr",
	}

	token, err := builder.Build(claims)
	checkErr(err)

	return token.String(), nil
}

func TokenGetEmail(token *jwt.Token) string {
	key := []byte("jwt_!test")

	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	checkErr(err)

	newToken, err := jwt.Parse(token.Bytes(), verifier)
	checkErr(err)

	err = verifier.Verify(newToken)
	checkErr(err)

	var u = userClaims{}

	err = json.Unmarshal(newToken.Claims(), &u)
	if err != nil {
		return ""
	}

	return u.Email
}

func StringToToken(strToken string) *jwt.Token {
	rawToken := []byte(strToken)

	key := []byte(`jwt_!test`)
	verifier, _ := jwt.NewVerifierHS(jwt.HS256, key)
	token, _ := jwt.Parse(rawToken, verifier)

	return token
}

func checkErr(err error) {
	if err != nil {
		log.Panicf("%v", err)
	}
}
