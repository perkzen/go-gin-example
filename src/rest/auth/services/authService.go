package services

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"rest-api/src/db/models"
)

type AuthService struct{}

func (userService AuthService) Create(user *models.User) error {
	err := mgm.Coll(user).Create(user)
	if err != nil {
		return err
	}
	return err
}

func (userService AuthService) FindOne(email string) (*models.User, error) {
	foundUser := &models.User{}
	err := mgm.Coll(foundUser).First(bson.M{"email": email}, foundUser)

	if err != nil {
		return nil, err
	}

	return foundUser, nil
}
