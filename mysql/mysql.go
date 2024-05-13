package mysql

import (
	"fmt"
	"project/model"
)

func GetPersonInfo(personID string) (model.Person, error) {
    
	db, err := OpenDB()
	if err != nil {
		fmt.Println("db not connected")
	}
    // Query to fetch person info
	var resp model.Person
    query := `select p.name , p2.number as phone_number, a.city , a.state, a.street1, a.street2, a.zip_code  from person p 
	left join phone p2 on p2.person_id = p.id 
	left join address_join aj on aj.person_id = p.id 
	left join address a on a.id = aj.id where p.id = ? `
    err1 := db.QueryRow(query, personID).Scan(&resp.Name, &resp.Phone_number, &resp.City, &resp.State, &resp.Street1, &resp.Street2, &resp.ZipCode)
    if err1 != nil {
        return resp, err
    }

    return resp, nil
}

func CreatePerson(input model.Person) (bool, error) {
    
	db, err := OpenDB()
	if err != nil {
		fmt.Println("db not connected")
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error starting transaction:", err)
	}
	personInsertQuery := "INSERT INTO person (name) VALUES (?)"
	personResult, err := db.Exec(personInsertQuery, input.Name)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error inserting into person table:", err)

	}

	// Retrieve the last inserted person ID
	personID, err := personResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println("Error getting last inserted person ID:", err)
		
	}

	// Insert into phone table
	phoneInsertQuery := "INSERT INTO phone (number, person_id) VALUES (?, ?)"
	_, err = db.Exec(phoneInsertQuery, input.Phone_number, personID)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error inserting into phone table:", err)
		
	}

	// Insert into address table
	addressInsertQuery := "INSERT INTO address (city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)"
	addressResult, err := db.Exec(addressInsertQuery, input.City, input.State, input.Street1, input.Street2, input.ZipCode)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error inserting into address table:", err)
		
	}

	// Retrieve the last inserted address ID
	addressID, err := addressResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println("Error getting last inserted address ID:", err)
		
	}

	// Insert into address_join table
	addressJoinInsertQuery := "INSERT INTO address_join (person_id, address_id) VALUES (?, ?)"
	_, err = db.Exec(addressJoinInsertQuery, personID, addressID)
	if err != nil {
		fmt.Println("Error inserting into address_join table:", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error committing transaction:", err)
	}
	return true, nil
	
}