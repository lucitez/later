package inits

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// driver

	_ "github.com/lib/pq"
)

// DB Creates global connection to postgres db using hardcoded values
func DB() *sql.DB {
	var (
		connectionName = mustGetenv("DB_HOST")
		user           = mustGetenv("DB_USER")
		dbName         = os.Getenv("DB_NAME") // NOTE: dbName may be empty
		password       = os.Getenv("DB_PASS") // NOTE: password may be empty
		socket         = os.Getenv("DB_SOCKET")
	)

	// connection string format: user=USER password=PASSWORD host=/cloudsql/PROJECT_ID:REGION_ID:INSTANCE_ID/[ dbname=DB_NAME]
	dbURI := fmt.Sprintf("user=%s password=%s host=%s%s dbname=%s", user, password, socket, connectionName, dbName)
	conn, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to ping db: %v", err)
	}

	return conn
}

// TestDB Creates global connection to postgres db using hardcoded values
func TestDB() *sql.DB {
	var (
		connectionName = mustGetenv("TEST_DB_HOST")
		user           = mustGetenv("TEST_DB_USER")
		dbName         = os.Getenv("TEST_DB_NAME") // NOTE: dbName may be empty
		password       = os.Getenv("TEST_DB_PASS") // NOTE: password may be empty
		socket         = os.Getenv("TEST_DB_SOCKET")
	)

	// connection string format: user=USER password=PASSWORD host=/cloudsql/PROJECT_ID:REGION_ID:INSTANCE_ID/[ dbname=DB_NAME]
	dbURI := fmt.Sprintf("user=%s password=%s host=%s%s dbname=%s", user, password, socket, connectionName, dbName)
	conn, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to ping db: %v", err)
	}

	return conn
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
