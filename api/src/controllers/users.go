package controllers

import "net/http"

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
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
