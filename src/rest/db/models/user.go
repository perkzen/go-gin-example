package models

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email"`
	Password         string `json:"password"`
}
