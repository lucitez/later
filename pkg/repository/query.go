package repository

import (
	"fmt"
	"strconv"
	"strings"
)

// Query struct for constructing DB queries
type Query struct {
	TableName        string
	SelectStatements []Select
	WhereStatements  []Where
	JoinStatements   []Join
	OrderStatements  []Order
}

// Select struct for organizing SELECT part of db query
type Select struct {
	TableName  string
	ColumnName string
}

// Where struct for organizing where statements for a db query
type Where struct {
	TableName  string
	ColumnName string
	Argument   *string
}

// Join struct for organizing joins for a db query
type Join struct {
	TableName      string
	ColumnName     string
	JoinTableName  string
	JoinColumnName string
}

// Order struct for organizing order statements in a db query
type Order struct {
	TableName  string
	ColumnName string
	Order      string
}

/**
* SELECT <table_name>.<column_name>
* FROM <table_name>
* JOIN <table_name2> ON <table_name2>.<column_name> = <table_name>.<column_name>
* WHERE <table_name>.<column_name> = $<counter>
 */

func generateSelectStatement(query *Query) string {
	statement := "SELECT "

	for _, where := range query.SelectStatements {
		statement += where.TableName + "." + where.ColumnName + ", "
	}

	return strings.TrimRight(statement, ", ")
}

func generateFromStatement(query *Query) string {
	return "FROM " + query.TableName
}

func generateJoinStatement(query *Query) string {
	statement := ""
	for _, joinStatement := range query.JoinStatements {
		statement +=
			"JOIN " + joinStatement.JoinTableName +
				" ON " + joinStatement.TableName + "." + joinStatement.ColumnName +
				" = " + joinStatement.JoinTableName + "." + joinStatement.JoinColumnName +
				" "
	}

	return strings.TrimRight(statement, " ")
}

func generateWhereStatement(query *Query) string {
	counter := 1
	statement := ""

	if len(query.WhereStatements) > 0 {
		statement += "WHERE "
	}

	for _, whereStatement := range query.WhereStatements {
		statement += whereStatement.TableName + "." + whereStatement.ColumnName
		if whereStatement.Argument != nil {
			statement += " = $" + strconv.Itoa(counter)
		}
		statement += " AND "
		counter++
	}

	return strings.TrimRight(statement, "AND ")
}

func generateOrderByStatement(query *Query) string {
	statement := ""

	if len(query.OrderStatements) > 0 {
		statement += " ORDER BY"
	}

	for _, orderStatement := range query.OrderStatements {
		statement += orderStatement.TableName + "." + orderStatement.ColumnName + " " + orderStatement.Order + ", "
	}

	return strings.TrimRight(statement, ", ")
}

// GenerateQuery creates an sql query string out of the Query struct
func (query *Query) GenerateQuery() string {
	selectStatement := generateSelectStatement(query)
	fromStatement := generateFromStatement(query)
	joinStatement := generateJoinStatement(query)
	whereStatement := generateWhereStatement(query)
	orderByStatement := generateOrderByStatement(query)

	statement := fmt.Sprintf("%s %s", selectStatement, fromStatement)
	if joinStatement != "" {
		statement = statement + " " + joinStatement
	}
	if whereStatement != "" {
		statement = statement + " " + whereStatement
	}
	if orderByStatement != "" {
		statement = statement + " " + orderByStatement
	}

	fmt.Println(statement)

	return statement + ";"
}

// GenerateArguments creates a list of interface from the Query class
func (query *Query) GenerateArguments() []interface{} {
	args := make([]interface{}, len(query.WhereStatements))
	for i, whereStatement := range query.WhereStatements {
		args[i] = whereStatement.Argument
	}

	return args
}
