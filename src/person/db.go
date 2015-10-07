package person

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Database variable useable from anywhere within our application
var (
	db *sql.DB
)

// Runs when the Go file is initialized
func init() {
	// Prepare our database connection string
	d := fmt.Sprintf("%s:%s@tcp4(%s:%s)/%s",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.DB,
	)
	// Attempt to establish a new SQL driver connection (does not actually connect to the SQL server)
	var err error
	db, err = sql.Open("mysql", d)
	// Something went wrong when trying to establish a driver connection to SQL
	if err != nil {
		panic(err)
	}
	// Ping does not open a connection, will only validate DSN data
	// Something went wrong with our DSN data sent to the MySQL server
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
