package person

import (
	"testing"
)

var p = Person{
	FirstName:    "Bob",
	LastName:     "Smith",
	EmailAddress: "notarealaddress@fakermcfakepants.com",
	Phone:        "555-555-5555",
	Weight:       180,
}

func TestCreate(t *testing.T) {
	t.Log("Attempting to create a new person")
	if err := p.Create(); err != nil {
		t.Errorf("Failed to create a new person!  %s", err)
	}
}

func TestDelete(t *testing.T) {
	t.Log("Attempting to delete an existing person")
	if err := p.Delete(); err != nil {
		t.Errorf("Failed to delete an existing person!  %s", err)
	}
}
