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

// GenerateUpdateStatement generates the select statement and arguments given an update body
// NOTE will NOT update fields to null.
// NOTE id field must be named "ID" in the update body
// NOTE if update body contains any structs other than a NullX variant, you're gonna have a bad time
func GenerateUpdateStatement(tableName string, updateBody interface{}) (string, []interface{}) {
	statement := "UPDATE " + tableName + " SET "
	arguments := []interface{}{}

	val := reflect.ValueOf(updateBody)
	typ := reflect.TypeOf(updateBody)

	counter := 1

	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		fieldName := stringutil.ToSnakeCase(typ.Field(i).Name)
		fieldValue := field.Interface()

		// If struct, we know it is a NullX variant. Find valid field
		if field.Kind() == reflect.Struct {
			valid := field.FieldByName("Valid").Bool()
			if valid {
				statement += fieldName + "=$" + strconv.Itoa(counter) + ", "
				arguments = append(arguments, fieldValue)
				counter++
			}
		} else {
			arguments = append(arguments, fieldValue)
			counter++
		}
	}

	statement = strings.TrimRight(statement, ", ")
	statement += " WHERE id = $1;"

	return statement, arguments
}
