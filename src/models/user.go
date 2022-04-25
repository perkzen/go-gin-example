package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kamva/mgm/v3"
	"rest-api/src/utils"
	"time"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
}

func (user *User) GenerateJwtToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	secretKey := utils.EnvVar("SECRET_KEY", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
