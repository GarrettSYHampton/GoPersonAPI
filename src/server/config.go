package main

import (
	"os"
	"strconv"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var config Config

func init() {

	if os.Getenv("WEBSERVERHOST") == "" {
		panic("No \"WEBSERVERHOST\" environment variable has been configured")
	}
	config.Host = os.Getenv("WEBSERVERHOST")

	var err error
	config.Port, err = strconv.Atoi(os.Getenv("WEBSERVERPORT"))
	if err != nil {
		panic("Environment variable \"WEBSERVERPORT\" needs to be an integer")
	}
}
