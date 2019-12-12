package model

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"golang.org/x/crypto/bcrypt"
)

//User do not know about how ot saves itself in db for that we need repository
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

//Validate is a method to validate inputs
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Email, validation.Required, validation.Length(6, 100)),
	)
}

//BeforeCreating where hasing happens
func (u *User) BeforeCreating() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
