package util

import (
	"database/sql"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// RepositoryTestUtil contains helper methods for running repository tests
type RepositoryTestUtil struct {
	DB     *sql.DB
	Assert *assert.Assertions
}

// TruncateTables clears all the tables so we don't violate duplicate indices
func (util *RepositoryTestUtil) TruncateTables(tableNames []string) {
	for _, tableName := range tableNames {
		util.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName))
	}
}
