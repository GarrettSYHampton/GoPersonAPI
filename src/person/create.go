package person

import (
	"log"
)

func Create(obj *Model) error {
	// Prepare a list of our SQL parameters
	params := make([]interface{}, 5)
	params[0] = obj.FirstName
	params[1] = obj.LastName
	params[2] = obj.EmailAddress
	params[3] = obj.Phone
	params[4] = obj.Weight
	// Attempt to run the SQL query against the database
	result, err := db.Exec(`
		INSERT INTO person (
			firstname,
			lastname,
			emailaddress,
			phone,
			weight
		)
		VALUES (
			?,
			?,
			?,
			?,
			?
		)`, params...)
	// If something failed when trying to run the SQL query
	if err != nil {
		log.Print(err)
		return err
	}
	// Attempt to grab the ID of the newly created object
	newID, err := result.LastInsertId()
	// if there was an error getting the newly created ID
	if err != nil {
		log.Print(err)
		return err
	}
	// No error when grabbing the newly created ID
	obj.ID = newID
	return nil
}
