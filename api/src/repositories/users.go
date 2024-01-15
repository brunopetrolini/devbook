package repositories

import (
	"database/sql"
	"devbook/src/models"
)

type users struct {
	db *sql.DB
}

// UsersRepository returns a users repository
func UsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Insert creates a new user in the database
func (u users) Insert(user models.User) (uint64, error) {
	statement, error := u.db.Prepare("INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastInsertedID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedID), nil
}
