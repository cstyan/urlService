package main

import (
	"fmt"
	"./urlsearch"
	// "github.com/jessevdk/go-flags"
	// "os"
)

// type options struct {
// 	ListenPort int  `short:"l" long:"listen-port" value-name:"<port number>" description:"Port to listen on for REST requests." required:"true"`
// 	Help       bool `short:"h" long:"help" description:"Show usage information."`
// }

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
	dataStore := urlsearch.NewLocalDataStore()
	dataStore.Upload("google.com,google.ca,youtube.com,amazon.com,amazon.com/a_thing")
	fmt.Println(dataStore.Query("amazon.com"))
	fmt.Println(dataStore.Query("amazon.com/a_thing"))
	fmt.Println(dataStore.Query("amazon.com/a-thing"))
}