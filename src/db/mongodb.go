package db

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"rest-api/src/utils"
)

func Init() {
	connectionString := utils.EnvVar("DB_CONNECTION_STRING", "")
	dbName := utils.EnvVar("DB_NAME", "")
	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal("Connection failed")
	}
}
