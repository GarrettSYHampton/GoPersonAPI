package main

import (
	"log"
	"net/http"
	"person"
)

// Request is a something...
type Request struct {
	Obj person.Model `json:"person"`
}

// Response is a something...
type Response struct {
	Error  error        `json:"error"`
	Object person.Model `json:"object,omitempty"`
}

// HTTPError logs an error and writes an error to the given
// http.ResponseWriter.
func HTTPError(w http.ResponseWriter, err error, status int) {
	log.Print(err)
	http.Error(w, err.Error(), status)
}
