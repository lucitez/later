package inits

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// driver

	_ "github.com/lib/pq"
	"github.com/lucitez/later/pkg/util/env"
)

// DB Creates global connection to postgres db using hardcoded values
func DB() *sql.DB {
	var (
		host     = env.MustGetenv("DB_HOST")
		user     = env.MustGetenv("DB_USER")
		dbName   = os.Getenv("DB_NAME") // NOTE: dbName may be empty
		password = os.Getenv("DB_PASS") // NOTE: password may be empty
	)

	// connection string format: user=USER password=PASSWORD host=/cloudsql/PROJECT_ID:REGION_ID:INSTANCE_ID/[ dbname=DB_NAME]
	dbURI := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", user, password, host, dbName)

	conn, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to ping db: %v", err)
	}

	fmt.Println("Initialized db.")

	return conn
}

// TestDB Creates global connection to postgres db using hardcoded values
func TestDB() *sql.DB {
	var (
		host     = env.MustGetenv("TEST_DB_HOST")
		user     = env.MustGetenv("TEST_DB_USER")
		dbName   = os.Getenv("TEST_DB_NAME") // NOTE: dbName may be empty
		password = os.Getenv("TEST_DB_PASS") // NOTE: password may be empty
	)

	// connection string format: user=USER password=PASSWORD host=/cloudsql/PROJECT_ID:REGION_ID:INSTANCE_ID/[ dbname=DB_NAME]
	dbURI := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", user, password, host, dbName)

	conn, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to ping db: %v", err)
	}

	return conn
}
