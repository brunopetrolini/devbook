package controllers

import (
	"devbook/src/application/responses"
	"devbook/src/domain/models"
	"devbook/src/infra/adapters/authentication"
	"devbook/src/infra/database"
	"devbook/src/infra/repositories"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("register"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	ID, error := repository.Insert(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	response := map[string]uint64{"id": ID}
	responses.JSON(w, http.StatusCreated, response)
}

// GetUsers returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	users, error := repository.GetUsers(nameOrNickname)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// FindUser returns a user by ID
func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	user, error := repository.FindUser(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	authenticatedUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userID != authenticatedUserID {
		responses.Error(w, http.StatusForbidden, errors.New("it's not possible to edit a user that is not yours"))
		return
	}

	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("update"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	if error = repository.UpdateUser(userID, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	authenticatedUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userID != authenticatedUserID {
		responses.Error(w, http.StatusForbidden, errors.New("it's not possible to delete a user that is not yours"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	if error = repository.DeleteUser(userID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser follows a user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	authenticatedUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	params := mux.Vars(r)

	followerID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerID == authenticatedUserID {
		responses.Error(w, http.StatusForbidden, errors.New("it's not possible to follow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.FollowersRepository(db)
	if error = repository.FollowUser(authenticatedUserID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser unfollows a user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	authenticatedUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)

	followerID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerID == authenticatedUserID {
		responses.Error(w, http.StatusForbidden, errors.New("it's not possible to unfollow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.FollowersRepository(db)
	if error = repository.UnfollowUser(authenticatedUserID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
