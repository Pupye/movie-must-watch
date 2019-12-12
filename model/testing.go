package model

import "testing"

//TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "userExample.org",
		Password: "password",
	}
}
