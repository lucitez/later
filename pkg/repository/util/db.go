package util

import (
	"database/sql"
	"fmt"

	// driver
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "later"
	password = ""
	dbname   = "later"

	testHost     = "localhost"
	testPort     = 5432
	testUser     = "later_test"
	testPassword = ""
	testDbname   = "later_test"
)

// InitDB Creates global connection to postgres db using hardcoded values
func InitDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// InitTestDB Creates global connection to postgres db using hardcoded values
func InitTestDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		testHost,
		testPort,
		testUser,
		testPassword,
		testDbname)

	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
