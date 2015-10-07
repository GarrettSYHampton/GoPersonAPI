package person

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

var (
	db   *sql.DB
	Conf Config
)

type Config struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port int    `json:"port"`
	DB   string `json:"db"`
}

type Person struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	EmailAddress string `json:"emailaddress"`
	Phone        string `json:"phone"`
	Weight       int    `json:"weight"`
}

func init() {

	port, err := strconv.Atoi(os.Getenv("MYSQLPORT"))
	if err != nil {
		panic(err)
	}

	Conf = Config{
		User: os.Getenv("MYSQLUSER"),
		Pass: os.Getenv("MYSQLPASS"),
		Host: os.Getenv("MYSQLHOST"),
		Port: port,
		DB:   os.Getenv("MYSQLDB"),
	}

	d := fmt.Sprintf("%s:%s@tcp4(%s:%s)/%s",
		Conf.User,
		Conf.Pass,
		Conf.Host,
		strconv.Itoa(Conf.Port),
		Conf.DB,
	)
	// Attempt to establish a new SQL driver connection (does not actually connect to the SQL server)
	db, err = sql.Open("mysql", d)
	if err != nil {
		panic(err)
	}
	// Ping does not open a connection, will only validate DSN data
	if err := db.Ping(); err != nil {
		panic(err)
	}
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
	if err != nil {
		return err
	}
	return nil
}
