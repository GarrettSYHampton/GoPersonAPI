# Go - "Person" API Demo
Project is supposed to be a web server that takes in JSON, creates or deletes data in MySQL, and responds with JSON.

## Dependencies
* go get github.com/gorilla/mux
* go get github.com/go-sql-driver/mysql

## Requirements
There are a few requirements handed to us from the "powers above":

* Configurations must be handled through environment variables
* No single quotes anywhere, only double quotes
* Application must not start if a database connection can not be established
* No "panics" anywhere
* Database CRUD actions and web server must be in separate packages
* Database must be MySQL

## Environment Variables
* Go
    * GOPATH
    * GOBIN
* Web Server
    * WEBSERVERHOST
    * WEBSERVERPORT
* MySQL Database
    * MYSQLHOST
    * MYSQLPORT
    * MYSQLUSER
    * MYSQLPASS
    * MYSQLDB

## Setup Instructions
* Setup local environment variables
* Run MySQL create.sql
* go install server
* . /bin/server
