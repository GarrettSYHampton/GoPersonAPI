package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, strconv.Itoa(config.Port)), router))
}
