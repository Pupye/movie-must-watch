package store_test

import "testing"

import "os"

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	//TODO compose using docker testing database
	databaseURL = os.Getenv("DARABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=almus password=pwd1234567 host=localhost dbname=mydb port=5433 sslmode=disable"
	}

	os.Exit(m.Run())
}
