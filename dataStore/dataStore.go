package dataStore

import (
	"log"
	"strings"
)

type DataStore interface {
	Clear() error
	Query(url string) bool
	// for now lets assume urls will be uploaded to us
	// via a comma separated list as one string
	Upload(urls string) bool
	// data stores that are accessing databases should
	// pull authentication from env vars instead of
	// by passing params to this function
	Initialize() error
	String()
}

type LocalDataStore struct {
	storage map[string]bool
}

func (data LocalDataStore) Clear() error {
	data.storage = make(map[string]bool)
	return nil
}

// we just want to know "does this url exist in the list"
// from the description we don't actually need to know the value
// any url in the list has the property the consumer wants to know about
func (data LocalDataStore) Query(url string) bool {
	log.Println("checking data store for url: ", url)
	_, ok := data.storage[url]
	return ok
}

// lets not care about the case where we're sent duplicate urls
// or urls the store already knows about for now
// it's not entirely clear if that's a case we actually care about
func (data LocalDataStore) Upload(urls string) bool {
	eachUrl := strings.Split(urls, ",")
	for _, url := range eachUrl {
		log.Println("adding url ", url, "to data store.")
		data.storage[url] = true
	}
	// remote data stores (redis) may fail for some reason?
	return true
}

func (data LocalDataStore) Initialize() error {
	// don't really need this function for local data store
	return nil
}

// for convenience/testing
func (data LocalDataStore) String() {
	for key, _ := range data.storage {
		log.Println(key)
	}
}

func NewLocalDataStore() LocalDataStore {
	var data LocalDataStore
	data.storage = make(map[string]bool)
	return data
}
