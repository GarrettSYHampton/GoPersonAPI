package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("WEBSERVERHOST"), os.Getenv("WEBSERVERPORT")), router))
}
