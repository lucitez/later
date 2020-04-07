package util

import (
	"later/pkg/util/stringutil"
	"reflect"
	"strconv"
	"strings"
)

// GenerateArguments creates a list of interface from the Query class
func GenerateArguments(arguments ...interface{}) []interface{} {
	args := []interface{}{}
	for _, argument := range arguments {
		reflectVal := reflect.ValueOf(argument)
		if reflectVal.Kind() == reflect.Ptr {
			if !reflectVal.IsNil() {
				args = append(args, argument)
			}
		} else {
			args = append(args, argument)
		}
	}

	return args
}

// GenerateInsertStatement creates a generic select statement maintaining struct field order
func GenerateInsertStatement(i interface{}, tableName string) string {
	t := reflect.TypeOf(i)
	numFields := t.NumField()

	statement := "INSERT INTO " + tableName + " ("

	for ind := 0; ind < numFields; ind++ {
		fieldName := stringutil.ToSnakeCase(t.Field(ind).Name)

		statement += fieldName + ","
	}

	statement = strings.TrimRight(statement, ",")
	statement += ") VALUES ("
	for ind := 1; ind <= numFields; ind++ {
		statement += "$" + strconv.Itoa(ind) + ","
	}

	statement = strings.TrimRight(statement, ",")
	statement += ");"

	return statement
}

// GenerateSelectStatement creates a generic select statement maintaining struct field order
func GenerateSelectStatement(i interface{}, tableName string) string {
	t := reflect.TypeOf(i)
	numFields := t.NumField()

	statement := "SELECT "

	for ind := 0; ind < numFields; ind++ {
		fieldName := stringutil.ToSnakeCase(t.Field(ind).Name)

		statement += tableName + "." + fieldName + ","
	}

	statement = strings.TrimRight(statement, ",")
	statement += " FROM " + tableName
	return statement
}
