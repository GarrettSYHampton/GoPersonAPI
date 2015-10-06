package person

import (
	"log"
)

func Delete(obj Model) error {
	// Prepare a list of our SQL parameters
	params := make([]interface{}, 1)
	params[0] = obj.ID
	// Attempt to run the SQL query against the database
	_, err := db.Exec(`
		DELETE
        FROM person
        WHERE
        id = ?
    `, params...)
	// If something failed when trying to run the SQL query
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
