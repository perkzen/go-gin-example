package main

import (
	"log"
	"rest-api/routes"
)

func main() {
	router := routes.InitRoute()
	err := router.Run()
	if err != nil {
		log.Fatal("error")
	}
	print("hello")
}
