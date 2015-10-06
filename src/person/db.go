package person

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

// Database variable useable from anywhere within our application
var (
	db *sql.DB
)

// Runs when the Go file is initialized
func init() {
	// Prepare our database connection string
	d := fmt.Sprintf("%s:%s@tcp4(%s:%s)/%s",
		os.Getenv("MYSQLUSER"),
		os.Getenv("MYSQLPASS"),
		os.Getenv("MYSQLHOST"),
		os.Getenv("MYSQLPORT"),
		os.Getenv("MYSQLDB"),
	)
	// Attempt to establish a new SQL driver connection (does not actually connect to the SQL server)
	database, err := sql.Open("mysql", d)
	// Something went wrong when trying to establish a driver connection to SQL
	if err != nil {
		panic(err)
	}
	// Ping does not open a connection, will only validate DSN data
	err = database.Ping()
	// Something went wrong with our DSN data sent to the MySQL server
	if err != nil {
		panic(err)
	}
	db = database
	// We successfully can communicate with the MySQL database server
	log.Print("Database connected successfully!")
}
