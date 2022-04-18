package models

import (
	"github.com/goonode/mogo"
)

type User struct {
	mogo.DocumentModel `bson:",inline" coll:"users"`
	Name               string
	Password           string
}

func init() {
	mogo.ModelRegistry.Register(User{})
}
