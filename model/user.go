package model

import "golang.org/x/crypto/bcrypt"

//User do not know about how ot saves itself in db for that we need repository
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
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
