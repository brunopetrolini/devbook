package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"encoding/json"
	"io"
	"net/http"
)

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, error := io.ReadAll(r.Body)
	if error != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	db, error := database.Connect()
	if error != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}

	repository := repositories.UsersRepository(db)
	ID, error := repository.Insert(user)
	if error != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]uint64{"id": ID}
	if error := json.NewEncoder(w).Encode(response); error != nil {
		http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
		return
	}
}

// GetUsers returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}

// FindUser returns a user by ID
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding a user by ID"))
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
