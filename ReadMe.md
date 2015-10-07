# Go - "Person" API Demo
Project is supposed to be a web server that takes in JSON, creates or deletes data in MySQL, and responds with JSON.

## Dependencies
* go get github.com/gorilla/mux
* go get github.com/go-sql-driver/mysql

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
