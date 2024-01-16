package models

import (
	"errors"
	"strings"
	"time"
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
	u.format()
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

	if step == "register" && u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)
}
