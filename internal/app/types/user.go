package types

import (
	"errors"
	"fmt"
	"net/mail"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func NewUser(firstName, lastName, email, password string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("First name cannot be empty")
	}
	if lastName == "" {
		return nil, errors.New("Last name cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	_, err := mail.ParseAddress(email)

	if err != nil {
		return nil, errors.New("invalid email")
	}

	return &User{
		gorm.Model{}, firstName, lastName, email, password,
	}, nil
}

func (u User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type CreateAccountInput struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
