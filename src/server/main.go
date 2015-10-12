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

	r := []string {"MYSQLUSER","MYSQLHOST","MYSQLPORT","MYSQLDB"}
	for i := range r {
		if os.Getenv(r[i]) == "" {
			log.Fatal("No \""+ r[i] +"\" environment variable has been configured")
			os.Exit(1)
		}
	}

	var err error
	var port int
	port, err = strconv.Atoi(os.Getenv("MYSQLPORT"))
	if err != nil {
		log.Fatal("Environment variable \"MYSQLPORT\" needs to be an integer")
		os.Exit(1)
	}

	err := person.Connect(os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASS"), os.Getenv("MYSQLHOST"), port, os.Getenv("MYSQLDB"))
	if err != nil {
		log.Fatal("Could not connect to database")
		os.Exit(1)
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, strconv.Itoa(config.Port)), router))
}
