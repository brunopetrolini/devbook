package database

import (
	"database/sql"
	"devbook/src/infra/config"

	_ "github.com/go-sql-driver/mysql"
)

// Connect opens a connection to the database
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.ConnectionString)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}
