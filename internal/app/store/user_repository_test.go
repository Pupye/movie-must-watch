package store_test

import "testing"

import "github.com/Pupye/movie-must-watch/model"

import "github.com/stretchr/testify/assert"

import "github.com/Pupye/movie-must-watch/internal/app/store"

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email:             "user@mail.org",
		EncryptedPassword: "asdasd",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.NotEqual(t, "", u.EncryptedPassword, "password should not be empty")
}
