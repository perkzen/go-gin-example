package services

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"rest-api/src/db/models"
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
