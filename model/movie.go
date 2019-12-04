package model

import "time"

type Movie struct {
	ID int
	Title string
	ReleaseDate time.Time
}