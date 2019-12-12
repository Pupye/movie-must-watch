package store

import "github.com/Pupye/movie-must-watch/model"

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create creates model in database and returns it
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	err := u.BeforeCreating()
	if err != nil {
		return nil, err
	}

	err = r.store.db.QueryRow(
		"INSERT INTO Users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID) // maps returing id into parameter

	if err != nil {
		return nil, err
	}

	return u, nil //inserted id
}

//FindByEmail find user by email and return its model
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1", email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword)

	if err != nil {
		return nil, err
	}

	return u, nil
}
