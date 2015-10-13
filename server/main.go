package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"person"
	"strconv"
)

func main() {

	r := []string {"MYSQLUSER","MYSQLHOST","MYSQLPORT","MYSQLDB","WEBSERVERHOST","WEBSERVERPORT" }
	for i := range r {
		if os.Getenv(r[i]) == "" {
			log.Fatalf("No \"%s\" environment variable has been configured", r[i])
		}
	}

	port, err := strconv.Atoi(os.Getenv("MYSQLPORT"))
	if err != nil {
		log.Fatal("Environment variable \"MYSQLPORT\" needs to be an integer")
	}

	err = person.Connect(os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASS"), os.Getenv("MYSQLHOST"), port, os.Getenv("MYSQLDB"))
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("WEBSERVERHOST"), os.Getenv("WEBSERVERPORT"), router))
}
