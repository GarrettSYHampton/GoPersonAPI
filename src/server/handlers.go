package main

import (
	"encoding/json"
	"net/http"
)

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
