package user

import (
	"errors"
	db "gin_gorm/gorm"
	"gin_gorm/util/auth"
	"github.com/cristalhq/jwt/v4"
	"log"
)

func getUser(token string) string {
	return getId(token)
}

func getId(strToken string) string {
	rawToken := []byte(strToken)

	key := []byte(`jwt_!test`)
	verifier, _ := jwt.NewVerifierHS(jwt.HS256, key)

	token, _ := jwt.Parse(rawToken, verifier)

	return auth.ParseToken(token)
}

func Login(u User) string {
	str, _ := auth.CreateToken(u.ID)
	return str
}

func validateSave(u User) (error, error) {
	name := u.Name
	pw := u.Password

	var errName error
	var errPw error
	if name == "" {
		errName = errors.New("empty name")
	}

	if pw == "" {
		errPw = errors.New("empty pw")
	}

	return errName, errPw
}

func save(u User) (uint, error) {
	err1, err2 := validateSave(u)

	if err1 != nil {
		return 0, err1
	}
	if err2 != nil {
		return 0, err2
	}

	// io get duple user-name
	var exists bool
	err3 := db.DB.Model(User{}).
		Select("count(*) > 0").
		Where("name = ?", u.Name).
		Find(&exists).
		Error
	if err3 != nil {
		log.Printf("err : fail count user-name // %v", err1)
		return 0, err3
	}

	if exists {
		log.Printf("err : already exists user-name\n")
		return 0, errors.New("err : already exists user-name")
	}

	// io save user
	u.Create()

	return u.ID, nil
}
