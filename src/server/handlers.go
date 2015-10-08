package main

import (
	"encoding/json"
	"net/http"
)

// Request is a something...
type Request struct {
	Obj person.Person `json:"person"`
}

// Response is a something...
type Response struct {
	Error  error         `json:"error"`
	Object person.Person `json:"person,omitempty"`
}

// HTTPError logs an error and writes an error to the given
// http.ResponseWriter.
func HTTPError(w http.ResponseWriter, err error, status int) {
	log.Print(err)
	http.Error(w, err.Error(), status)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var req Request
	var err error
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	if err = req.Obj.Create(); err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	output, err := json.Marshal(Response{
		Error:  err,
		Object: req.Obj,
	})

	if err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	w.Write(output)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var req Request
	var err error
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	if err = req.Obj.Delete(); err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	output, err := json.Marshal(Response{
		Error: err,
	})

	if err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	w.Write(output)
}
