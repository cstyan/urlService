package urlservice

import (
	"../dataStore"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Service struct {
	DataStore dataStore.DataStore
	Router    *mux.Router
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
