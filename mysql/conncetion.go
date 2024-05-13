package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	// MySQL connection string
	connectionString := "root:1234@tcp(localhost:3306)/cetec"

	// Open database connection
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Ping database to verify connection
	err = db.Ping()
	if err != nil {
		db.Close() // Close the database connection if ping fails
		return nil, err
	}

	return db, nil
}