package main

import "net/http"

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
