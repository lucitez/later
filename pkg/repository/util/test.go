package util

import (
	"database/sql"
	"fmt"
	"later/pkg/util/slice"
	"reflect"
	"testing"
)

// RepositoryTestUtil contains helper methods for running repository tests
type RepositoryTestUtil struct {
	DB *sql.DB
}

// TruncateTables clears all the tables so we don't violate duplicate indices
func (util *RepositoryTestUtil) TruncateTables(tableNames []string) {
	for _, tableName := range tableNames {
		util.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName))
	}
}

// AssertEquals will fail if the two interfaces do not satisfy reflect.DeepEqual()
func AssertEquals(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("Actual:\n%+v\n", actual)
		fmt.Printf("Expected:\n%+v\n", expected)
		t.Error("Actual was different from expected")
	}
}

// AssertContainsOne will fail if the provided slice does not contain the provided
// Item. Item must be passed as a pointer
func AssertContainsOne(t *testing.T, actual interface{}, item interface{}) {
	contains := false

	actualSlice := slice.InterfaceSlice(actual)

	for _, i := range actualSlice {
		if reflect.DeepEqual(i, item) {
			contains = true
			break
		}
	}

	if !contains {
		fmt.Printf("Actual:\n%+v\n", actual)
		fmt.Printf("Item:\n%+v\n", item)
		t.Error("Slice did not contain item")
	}
}
