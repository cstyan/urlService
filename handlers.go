package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// TODO: better names for these handlers

func (s *Service) getHandler(writer http.ResponseWriter, req *http.Request) {
	requestVars := mux.Vars(req)
	if s.DataStore.Query(requestVars["check_url"]) {
		writer.WriteHeader(200)
		writer.Write([]byte("true"))
		log.Println("200, url found in data store.")
		return
	}
	writer.WriteHeader(200)
	writer.Write([]byte("false"))
	log.Println("200, url not found in data store.")
}

func (s *Service) uploadHandler(writer http.ResponseWriter, req *http.Request) {
	// for now we're just handling plain text comma separated lists of urls
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print("400, could not parse request.")
		writer.WriteHeader(400)
		writer.Write([]byte("could not parse request."))
	}
	s.DataStore.Upload(string(data))
	writer.WriteHeader(200)
	writer.Write([]byte("uploaded"))
	log.Println("200, list of urls uploaded.")
}
