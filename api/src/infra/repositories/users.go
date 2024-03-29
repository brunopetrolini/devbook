package repositories

import (
	"database/sql"
	"devbook/src/domain/models"
	"fmt"
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

// GetUsers returns all users filtered by name or nickname
func (u users) GetUsers(nameOrNickname string) ([]models.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname) // %nameOrNickname%

	lines, error := u.db.Query("SELECT id, name, nickname, email, created_at FROM users WHERE name LIKE ? OR nickname LIKE ?", nameOrNickname, nameOrNickname)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if error = lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

// FindUser returns a user by its ID
func (u users) FindUser(userId uint64) (models.User, error) {
	line, error := u.db.Query("SELECT id, name, nickname, email, created_at FROM users WHERE id = ?", userId)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error = line.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

// UpdateUser updates a user in the database
func (u users) UpdateUser(userID uint64, user models.User) error {
	statement, error := u.db.Prepare("UPDATE users SET name = ?, nickname = ?, email = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nickname, user.Email, userID); error != nil {
		return error
	}

	return nil
}

// DeleteUser deletes a user from the database
func (u users) DeleteUser(userID uint64) error {
	statement, error := u.db.Prepare("DELETE FROM users WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID); error != nil {
		return error
	}

	return nil
}

// SearchByEmail searches a user by its email
func (u users) SearchByEmail(email string) (models.User, error) {
	line, error := u.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error = line.Scan(&user.ID, &user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}
