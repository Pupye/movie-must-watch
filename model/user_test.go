package model_test

import (
	"testing"

	"github.com/Pupye/movie-must-watch/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreating())
	assert.NotEmpty(t, u.EncryptedPassword)
}
