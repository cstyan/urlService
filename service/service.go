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
	malicious map[string]bool
}

func (s *Service) InitService(mainDataStore dataStore.DataStore) {
	s.malicious = make(map[string]bool)
	s.malicious["whitelist"] = false
	s.malicious["blacklist"] = true
	// is there an easier way to replace the data store type
	// TODO: error check here for other data store types
	s.DataStore = mainDataStore
	s.newRouter()
}

func (s *Service) Run(listenAddress string) {
	log.Fatal(http.ListenAndServe(listenAddress, s.Router))
}
