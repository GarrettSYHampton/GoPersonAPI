package person

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

var db *sql.DB

type Person struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	EmailAddress string `json:"emailaddress"`
	Phone        string `json:"phone"`
	Weight       int    `json:"weight"`
}

func Connect(user string, pass string, host string, port int, schema string) error {
	d := fmt.Sprintf("%s:%s@tcp4(%s:%s)/%s",
		user,
		pass,
		host,
		strconv.Itoa(port),
		schema,
	)
	// Attempt to establish a new SQL driver connection (does not actually connect to the SQL server)
	var err error
	db, err = sql.Open("mysql", d)
	if err != nil {
		return err
	}
	// Ping does not open a connection, will only validate DSN data
	return db.Ping()
}

func (p *Person) Create() error {
	result, err := db.Exec(`
    		INSERT INTO person (
    			firstname,
    			lastname,
    			emailaddress,
    			phone,
    			weight
    		)
    		VALUES (
    			?, ?, ?, ?, ?
    		)`,
		p.FirstName, p.LastName, p.EmailAddress, p.Phone, p.Weight)

	if err != nil {
		return err
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = newID
	return nil
}

func (p Person) Delete() error {
	_, err := db.Exec(`
		DELETE
        FROM person
        WHERE
        id = ?
    `, p.ID)
	return err
}
