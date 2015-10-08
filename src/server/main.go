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

	if os.Getenv("MYSQLUSER") == "" {
		panic("No \"MYSQLUSER\" environment variable has been configured")
	}

	if os.Getenv("MYSQLHOST") == "" {
		panic("No \"MYSQLHOST\" environment variable has been configured")
	}

	if os.Getenv("MYSQLPORT") == "" {
		panic("No \"MYSQLPORT\" environment variable has been configured")
	}
	var err error
	var port int
	port, err = strconv.Atoi(os.Getenv("MYSQLPORT"))
	if err != nil {
		panic("Environment variable \"MYSQLPORT\" needs to be an integer")
	}
	if os.Getenv("MYSQLDB") == "" {
		panic("No \"MYSQLDB\" environment variable has been configured")
	}

	err := person.Connect(os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASS"), os.Getenv("MYSQLHOST"), port, os.Getenv("MYSQLDB"))
	if err != nil {
		panic("Could not connect to database")
	}
	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, strconv.Itoa(config.Port)), router))
}
