package dataStore

import (
	"fmt"
	"log"
	"strings"
)

type DataStore interface {
	Clear() error
	Query(url string) (bool, error)
	// for now lets assume urls will be uploaded to us
	// via a comma separated list as one string
	Upload(urls string, malicious bool) error
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
func (data LocalDataStore) Query(url string) (bool, error) {
	log.Println("checking data store for url: ", url)
	val, ok := data.storage[url]
	if !ok {
		// false is the zero value for a bool
		return false, fmt.Errorf("url: %s not found in data store", url)
	}
	return val, nil
}

// lets not care about the case where we're sent duplicate urls
// or urls the store already knows about for now
// it's not entirely clear if that's a case we actually care about
func (data LocalDataStore) Upload(urls string, malicious bool) error {
	eachUrl := strings.Split(urls, ",")
	for _, url := range eachUrl {
		log.Println("adding url ", url, "to data store.")
		data.storage[url] = malicious
	}
	// remote data stores (redis) may fail for some reason?
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
