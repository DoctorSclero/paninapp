package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model. Fields are kept public to
// give the ORM library the ability to change
// the content. Although normal generation should be
// done by using the #NewUser(email, password string) method

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"unique;not null"`
	Hash  string `json:"-" gorm:"not null"`
}

func NewUser(email, password string) (*User, error) {
	// Bcrypt hashing of password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// Creating user
	user := &User{Email: email, Hash: string(hash)}
	return user, nil
}

// Checks provided password against the stored hash
// to ensure they match. Returns true if the passwords
// match, false otherwise.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}
