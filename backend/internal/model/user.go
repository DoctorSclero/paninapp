package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique;not null"`
	Hash  string `gorm:"not null"`
}

func NewUserFromEmailAndPassword(email, password string) (*User, error) {
	// Bcrypt hashing of password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// Creating user
	user := &User{Email: email, Hash: string(hash)}
	return user, nil
}
