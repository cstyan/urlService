package main

import (
	"./dataStore"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

type Service struct {
	DataStore dataStore.DataStore
	Router *mux.Router
}

func (s *Service) queryDataStore(url string) bool {
	return s.DataStore.Query(url)
}

func (s *Service) InitService() {
	// is there an easier way to replace the data store type
	// TODO: error check here for other data store types
	s.DataStore = dataStore.NewLocalDataStore()
	s.newRouter()
}

func (s *Service) Run(listenAddress string) {
	log.Fatal(http.ListenAndServe(listenAddress, s.Router))
}
