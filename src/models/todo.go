package models

import "github.com/kamva/mgm/v3"

type Todo struct {
	mgm.DefaultModel `bson:",inline"`
	Text             string `json:"text" bson:"text"`
	Completed        bool   `json:"completed" bson:"completed"`
}
