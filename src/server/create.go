package main

import (
	"encoding/json"
	"log"
	"net/http"
	"person"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// Define the request we are expecting
	req := struct {
		Obj person.Model `json:"person"`
	}{}
	// Parse data from request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Attempt to create new object in database
	err := person.Create(&req.Obj)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Define what we want our response to look like
	res := struct {
		Error  error        `json:"error"`
		Object person.Model `json:"object"`
	}{
		Error:  err,
		Object: req.Obj,
	}
	// Marshal the response into JSON
	output, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Print the response
	w.Write(output)
}
