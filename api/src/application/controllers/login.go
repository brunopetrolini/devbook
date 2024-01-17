package controllers

import (
	"devbook/src/application/responses"
	"devbook/src/domain/models"
	"devbook/src/infra/adapters/authentication"
	encrypter "devbook/src/infra/adapters/encrypt"
	"devbook/src/infra/database"
	"devbook/src/infra/repositories"
	"encoding/json"
	"io"
	"net/http"
)

// Login is responsible for authenticating the user
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	persistedUser, error := repository.SearchByEmail(user.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = encrypter.Compare(persistedUser.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, _ := authentication.GenerateToken(persistedUser.ID)
	response := map[string]string{
		"access_token": token,
	}

	responses.JSON(w, http.StatusOK, response)
}
