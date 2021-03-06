package urlservice

import (
	"github.com/gorilla/mux"
	"net/http"
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
	var routes = Routes{
		Route{"GET", "/urlinfo/v1/{check_url}", s.getHandler},
		Route{"POST", "/urlinfo/v1/blacklist", s.uploadHandler},
		Route{"POST", "/urlinfo/v1/whitelist", s.uploadHandler},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	s.Router = router
}
