package main

import (
	"log"
	// "github.com/gorilla/mux"
	// "net/http"
	// "github.com/jessevdk/go-flags"
	"os"
)

// type options struct {
// 	ListenPort int  `short:"l" long:"listen-port" value-name:"<port number>" description:"Port to listen on for REST requests." required:"true"`
// 	Help       bool `short:"h" long:"help" description:"Show usage information."`
// }

// TODO: get rid of this global, need an easier way to pass data store to handlers
var service Service

func main() {
	// var opts options
	// parser := flags.NewParser(&opts, 0)
	// _, err := parser.ParseArgs(os.Args[1:])
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	// if opts.Help {
	// 	parser.WriteHelp(os.Stdout)
	// 	os.Exit(0)
	// }
	log.SetOutput(os.Stdout)
	service.InitService()
	log.Println("Starting URLCheck service.")

	// seed with some fake data to start
	service.DataStore.Upload("google.com,google.ca,youtube.com,amazon.com,amazon.com/a_thing")
	
	service.Run(":8080")
}