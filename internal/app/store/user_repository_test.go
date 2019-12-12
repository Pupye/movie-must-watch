package store_test

import "testing"

import "github.com/Pupye/movie-must-watch/model"

import "github.com/stretchr/testify/assert"

import "github.com/Pupye/movie-must-watch/internal/app/store"

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	s.User().Create(&model.User{
		Email:             "findthis@gmail.com",
		EncryptedPassword: "asdasd",
	})

	u, err := s.User().FindByEmail("findthis@gmail.com")

	assert.NoError(t, err, "unexpected error")
	assert.NotNil(t, u)
	assert.Equal(t, "findthis@gmail.com", u.Email)
}
