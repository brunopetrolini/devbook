package models

import (
	encrypter "devbook/src/infra/adapters/encrypt"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user in the social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will validate and format the user
func (u *User) Prepare(step string) error {
	if error := u.validate(step); error != nil {
		return error
	}
	if error := u.format(step); error != nil {
		return error
	}
	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Nickname == "" {
		return errors.New("nickname is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if error := checkmail.ValidateFormat(u.Email); error != nil {
		return errors.New("invalid email")
	}

	if step == "register" && u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)

	if step == "register" {
		hash, error := encrypter.Hash(u.Password)
		if error != nil {
			return error
		}
		u.Password = string(hash)
	}

	return nil
}
