package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Route struct {
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

// lets assume someone may change the api at some point in the future
// and set things up to use subrouters from the start
// apiBaseUrl := "/urlinfo/v1/"
// router := mux.newRouter()
// subrouter := router.PathPrefix(apiBaseUrl).Subrouter()
// subrouter.HandleFunc("/{url}", getHandler)

func (s *Service) newRouter() {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Handler(route.HandlerFunc)
    }

    s.Router = router
}

// handler functions are in handlers.go
// TODO: make this not global?
var routes = Routes{
	Route{"GET", "/urlinfo/v1/{check_url}", service.getHandler},
	Route{"POST", "/urlinfo/v1/upload", service.uploadHandler},
}