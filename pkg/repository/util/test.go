package util

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

type RepositoryTestUtil struct {
	DB *sql.DB
}

func (util *RepositoryTestUtil) TruncateTables(tableNames []string) {
	for _, tableName := range tableNames {
		util.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName))
	}
}

func AssertEquals(t *testing.T, actual interface{}, expected interface{}) {
	fmt.Printf("Actual:\n%+v\n", actual)
	fmt.Printf("Expected:\n%+v\n", expected)
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Actual was different from expected")
	}
}
