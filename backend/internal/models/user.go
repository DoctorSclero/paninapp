package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

var ErrorBadPassword = errors.New("password criteria not met")
var ErrorBadEmail = errors.New("bad email")
var ErrorWrongPassword = errors.New("wrong password")

type User struct {
	Email string `json:"email"`
	Hash  string `json:"-"`
}

// Creates a new user instance from email and password.
// Email must be valid and non empty, password must be
// at least 8 characters long. Throws ErrorBadEmail if
// email failed validation, ErrorBadPassword if
// password failed validation, BCrypt errors if hashing
// fails.
func (u *User) CreateFromEmailAndPassword(email, password string) error {
	// Validate email format
	if err := validate.Var(email, "required,email"); err != nil {
		return ErrorBadEmail
	}

	// Validate password format (e.g., minimum 8 characters)
	if err := validate.Var(password, "required,min=8"); err != nil {
		return ErrorBadPassword
	}

	// Generate hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Store data
	u.Hash = string(hash)
	u.Email = email
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}
