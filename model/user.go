package model

//User do not know about how ot saves itself in db for that we need repository
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
}
