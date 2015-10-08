package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	Route{
		Name:        "Create",
		Method:      "POST",
		Pattern:     "/person/create",
		HandlerFunc: Create,
	},
	Route{
		Name:        "Delete",
		Method:      "DELETE",
		Pattern:     "/person/delete",
		HandlerFunc: Delete,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Host("*")
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = LogRequest(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(DefaultHandler(handler))
	}
	return router
}

func LogRequest(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func DefaultHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Handle API authentication
		if r.Method == "OPTIONS" {
			//handle preflight in here
		} else {
			// DEV: Allow cross domain access
			w.Header().Set("Access-Control-Allow-Origin", "*")
			// All of our responses should be in JSON
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			h.ServeHTTP(w, r)
		}
	}
}
