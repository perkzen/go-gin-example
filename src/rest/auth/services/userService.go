package services

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"rest-api/src/rest/db/models"
)

type UserService struct{}

func (userService UserService) Create(user *models.User) error {
	err := mgm.Coll(user).Create(user)
	if err != nil {
		return err
	}
	return err
}

func (userService UserService) GetAll() ([]models.User, error) {
	var results []models.User
	err := mgm.Coll(&models.User{}).SimpleFind(&results, bson.M{})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (userService UserService) FindOne(user *models.User) (*models.User, error) {
	foundUser := &models.User{}
	err := mgm.Coll(user).First(bson.M{"email": user.Email}, foundUser)

	if err != nil {
		return nil, err
	}

	return foundUser, nil
}
