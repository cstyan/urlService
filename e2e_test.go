package main_test

import (
	"log"
	"bytes"
	"io/ioutil"
	"testing"
	"os"
	"./service"
	"net/http/httptest"
	"net/http"
)


// TODO: get rid of this global, need an easier way to pass data store to handlers
var service urlservice.Service

// helpers
func clearData() {
	service.DataStore.Clear()
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    service.Router.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

// tests
func TestMain(m *testing.M) {
	// remove this line if you want to see logging from internal functions
	log.SetOutput(ioutil.Discard)
	service.InitService()

	// run all tests
	code := m.Run()
	os.Exit(code)
}

func TestGetNoUrl(t *testing.T) {
	clearData()

	req, _ := http.NewRequest("GET", "/urlinfo/v1/asdf.com", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "false" {
        t.Errorf("Expected GET request to not find url (false). Got %s", body)
    }
}

func TestGetHasUrl(t *testing.T) {
	clearData()
	
	// TODO: can we rely on this in a test? not when using a remote store
	service.DataStore.Upload("asdf.com")
	
	req, _ := http.NewRequest("GET", "/urlinfo/v1/asdf.com", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "true" {
        t.Errorf("Expected GET request to find url (true). Got %s", body)
    }
}

// TODO: upload currently only errors if parsing the request body fails
// is there anything failure case we're not thinking of?

func TestUploadSuccess(t *testing.T) {
	clearData()

	payload := []byte("asdf.com,aaa.com,bbb.com")
	req, _ := http.NewRequest("POST", "/urlinfo/v1/upload", bytes.NewBuffer(payload))
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "uploaded" {
        t.Errorf("Expected POST request to upload urls (uploaded). Got %s", body)
    }
}
