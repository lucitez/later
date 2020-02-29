package repository_test

import (
	"testing"

	"github.com/google/uuid"
	"later.co/pkg/repository"
)

var tableName = "test_table"
var testID, _ = uuid.NewRandom()
var testIDString = testID.String()

var testSelect = repository.Select{
	TableName:  tableName,
	ColumnName: "test_column"}

var testWhere = repository.Where{
	TableName:  tableName,
	ColumnName: "test_column",
	Argument:   &testIDString}

var testJoin = repository.Join{
	TableName:      tableName,
	ColumnName:     "test_column",
	JoinTableName:  "join_table",
	JoinColumnName: "join_column"}

func assertStringEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf(expected+"\nGOT: %s", actual)
	}
}

func TestQueryGenerateQuerySelect(t *testing.T) {

	var testQuery = repository.Query{
		TableName:        tableName,
		SelectStatements: []repository.Select{testSelect}}

	actual := testQuery.GenerateQuery()
	expected := "SELECT test_table.test_column FROM test_table;"
	assertStringEquals(t, actual, expected)
}

func TestQueryGenerateQueryWhere(t *testing.T) {

	var testQuery = repository.Query{
		TableName:        tableName,
		SelectStatements: []repository.Select{testSelect},
		WhereStatements:  []repository.Where{testWhere}}

	actual := testQuery.GenerateQuery()
	expected := "SELECT test_table.test_column FROM test_table WHERE test_table.test_column = $1;"
	assertStringEquals(t, actual, expected)
}

func TestQueryGenerateQueryJoin(t *testing.T) {

	argument := "anything"

	joinSelect := repository.Select{
		TableName:  "join_table",
		ColumnName: "join_column"}

	joinWhere := repository.Where{
		TableName:  "join_table",
		ColumnName: "join_column",
		Argument:   &argument}

	testQuery := repository.Query{
		TableName:        tableName,
		SelectStatements: []repository.Select{testSelect, joinSelect},
		WhereStatements:  []repository.Where{testWhere, joinWhere},
		JoinStatements:   []repository.Join{testJoin}}

	actual := testQuery.GenerateQuery()
	expected :=
		"SELECT test_table.test_column, join_table.join_column " +
			"FROM test_table " +
			"JOIN join_table ON test_table.test_column = join_table.join_column " +
			"WHERE test_table.test_column = $1 AND join_table.join_column = $2;"

	assertStringEquals(t, actual, expected)
}
