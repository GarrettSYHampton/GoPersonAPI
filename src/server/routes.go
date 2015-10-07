package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name: "Create",
		Method: "POST",
		Pattern: "/person/create",
		HandlerFunc: Create,
	},
	Route{
		Name: "Delete",
		Method: "DELETE",
		Pattern: "/person/delete",
		HandlerFunc: Delete,
	},
}
