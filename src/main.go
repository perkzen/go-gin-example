package main

import (
	"log"
	"rest-api/src/db"
	_ "rest-api/src/db"
	"rest-api/src/routes"
	"rest-api/src/utils"
)

func main() {
	db.Init()
	router := routes.InitRoute()
	port := utils.EnvVar("SERVER_PORT", ":8080")
	err := router.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
