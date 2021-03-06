package main

import (
	"./service"
	"./dataStore"
	"log"
	"os"
)

func main() {
	var service urlservice.Service

	log.SetOutput(os.Stdout)
	service.InitService(dataStore.NewRedisDataStore())
	log.Println("Starting URLCheck service.")

	// seed with some fake data to start
	service.DataStore.Upload("google.com,google.ca,youtube.com,amazon.com,amazon.com/a_thing", true)

	service.Run(":8080")
}
