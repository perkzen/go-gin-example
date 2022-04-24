package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kamva/mgm/v3"
	"rest-api/src/utils"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
}

func (user *User) GenerateJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": user.Email})
	secretKey := utils.EnvVar("SECRET:KEY", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
