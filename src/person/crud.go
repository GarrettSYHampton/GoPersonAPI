package person

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
