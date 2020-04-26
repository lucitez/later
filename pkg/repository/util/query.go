package util

import (
	"database/sql"
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

func ScanRowsInto(rows *sql.Rows, i interface{}) error {
	elem := reflect.ValueOf(i).Elem()

	dest := make([]interface{}, elem.NumField())

	for i := 0; i < elem.NumField(); i++ {
		valueField := elem.Field(i)
		dest[i] = valueField.Addr().Interface()
	}

	return rows.Scan(dest...)
}

// ScanRowInto scans the db row into the struct passed in by reference
func ScanRowInto(row *sql.Row, i interface{}) error {
	elem := reflect.ValueOf(i).Elem()

	dest := make([]interface{}, elem.NumField())

	for i := 0; i < elem.NumField(); i++ {
		valueField := elem.Field(i)
		dest[i] = valueField.Addr().Interface()
	}

	return row.Scan(dest...)
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

// GenerateInsertArguments accepts a pointer to a struct and returns its field values as a []interface{}
func GenerateInsertArguments(i interface{}) []interface{} {
	elem := reflect.ValueOf(i).Elem()
	args := []interface{}{}

	for i := 0; i < elem.NumField(); i++ {
		args = append(args, elem.Field(i).Interface())
	}

	return args
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
