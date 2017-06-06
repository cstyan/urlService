package main

import (
	"log"
	"./service"
	"os"
)

func main() {
	var service urlservice.Service
	
	log.SetOutput(os.Stdout)
	service.InitService()
	log.Println("Starting URLCheck service.")

	// seed with some fake data to start
	service.DataStore.Upload("google.com,google.ca,youtube.com,amazon.com,amazon.com/a_thing")

	service.Run(":8080")
}
