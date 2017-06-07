package urlservice

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//get whitelist or blacklist from url so we can pass to upload
func getLastSectionOfPath(path string) string {
	chunks := strings.Split(path, "/")
	return chunks[len(chunks)-1]
}

// TODO: better names for these handlers

func (s *Service) getHandler(writer http.ResponseWriter, req *http.Request) {
	requestVars := mux.Vars(req)
	url := requestVars["check_url"]
	// TODO: handle the var not being present
	value, err := s.DataStore.Query(url)

	if err != nil {
		responseString := fmt.Sprintf("%s was not found .\n", url)
		writer.WriteHeader(200)
		writer.Write([]byte(responseString))
		log.Printf("200, %s", responseString)
		return
	}

	if value {
		responseString := fmt.Sprintf("%s is malicious", url)
		writer.WriteHeader(200)
		writer.Write([]byte(responseString))
		log.Printf("200, %s", responseString)
		return
	}
	responseString := fmt.Sprintf("%s is not malicious", url)
	writer.WriteHeader(200)
	writer.Write([]byte(responseString))
	log.Printf("200, %s", responseString)
}

// TODO: upload currently only errors if parsing the request body fails
// is there anything failure case we're not thinking of?
func (s *Service) uploadHandler(writer http.ResponseWriter, req *http.Request) {
	last := getLastSectionOfPath(req.URL.Path)
	// for now we're just handling plain text comma separated lists of urls
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print("400, could not parse request.")
		writer.WriteHeader(400)
		writer.Write([]byte("could not parse request."))
	}
	log.Println("attempting to upload urls: ", string(data))
	s.DataStore.Upload(string(data), s.malicious[last])
	writer.WriteHeader(200)
	writer.Write([]byte("uploaded"))
	log.Println("200, list of urls uploaded.")
}
