package person

import (
	"os"
	"strconv"
	"testing"
)

func TestCRUD(t *testing.T) {
	var err error
	var port int
	port, err = strconv.Atoi(os.Getenv("MYSQLPORT"))
	if err != nil {
		t.Errorf("Could not determine MySQL port from environment variables! %s", err)
		return
	}

	if err = Connect(os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASS"), os.Getenv("MYSQLHOST"), port, os.Getenv("MYSQLDB")); err != nil {
		t.Errorf("Error connecting to MySQL Database!  %s", err)
		return
	}

	p := Person{
		FirstName:    "Bob",
		LastName:     "Smith",
		EmailAddress: "notarealaddress@fakermcfakepants.com",
		Phone:        "555-555-5555",
		Weight:       180,
	}

	t.Log("Attempting to create a new person")
	if err = p.Create(); err != nil {
		t.Errorf("Failed to create a new person!  %s", err)
	}

	t.Log("Attempting to delete an existing person")
	if err := p.Delete(); err != nil {
		t.Errorf("Failed to delete an existing person!  %s", err)
	}
}
