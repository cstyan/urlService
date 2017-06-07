package dataStore

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var dataStore = NewLocalDataStore()

// helpers
func clearData() {
	dataStore.storage = make(map[string]bool)
}

// tests
func TestMain(m *testing.M) {
	// remove this line if you want to see logging from internal functions
	log.SetOutput(ioutil.Discard)

	// run all tests
	code := m.Run()
	os.Exit(code)
}

func TestClearLocalDataStore(t *testing.T) {
	dataStore.Clear()
	if len(dataStore.storage) != 0 {
		t.Error("Clear did not clear all data from localDataStore.")
	}
}

func TestQueryHasUrlLocalDataStore(t *testing.T) {
	clearData()

	dataStore.storage["asdf.com"] = true
	_, err := dataStore.Query("asdf.com")

	if err != nil {
		t.Error("Added URL to storage, but Query for that URL returned an error.")
	}
}

func TestQueryNoUrlLocalDataStore(t *testing.T) {
	clearData()

	_, err := dataStore.Query("asdf.com")
	if err == nil {
		t.Error("Local data store is empty, but Query for a URL returned true.")
	}
}

// upload doesn't really fail for local data store?
func TestUploadTrueLocalDataStore(t *testing.T) {
	clearData()

	if !dataStore.Upload("asdf.com,aaa.com", true) {
		t.Error("Upload returned false, this should never happen at the moment.")
	}
	if dataStore.storage["asdf.com"] != true || dataStore.storage["aaa.com"] != true {
		t.Error("Upload did not work properly, one or more of the specified URLs is not in the storage")
	}
}

// upload doesn't really fail for local data store?
func TestUploadFalseLocalDataStore(t *testing.T) {
	clearData()

	if !dataStore.Upload("asdf.com,aaa.com", false) {
		t.Error("Upload returned false, this should never happen at the moment.")
	}
	if dataStore.storage["asdf.com"] != false || dataStore.storage["aaa.com"] != false {
		t.Error("Upload did not work properly, one or more of the specified URLs is not in the storage")
	}
}
