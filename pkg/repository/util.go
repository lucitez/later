package repository

import (
	"database/sql"
	"strconv"
	"strings"
)

// ConstructWhens turns a list of database column names and turns them into when statements with injection
// ConstructWhens([]string{'foo', 'bar'}, 2) => " AND foo = $2 AND bar = $3"
func ConstructWhens(whens []string, startingIndex int) string {
	counter := startingIndex
	retString := ""

	for _, when := range whens {
		retString = retString + " AND " + when + " = $" + strconv.Itoa(counter)
		counter++
	}

	return strings.Trim(retString, " ")
}

// QueryRowsWithArguments takes a database connection, sql statement (with injections), and a list of arguments
// This function converts arguments ([]string) and converts them to []interface so that we can spread them to use
// In the DB.Query call
func QueryRowsWithArguments(db *sql.DB, statement string, arguments []string) (*sql.Rows, error) {
	args := make([]interface{}, len(arguments))
	for i, arg := range arguments {
		args[i] = arg
	}

	return db.Query(statement, args...)
}
