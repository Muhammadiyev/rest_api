package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=rest_api dbname=restapi_dev password=qwerty1234! sslmode=disable"
	}

	os.Exit(m.Run())
}
