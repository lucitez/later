package errs

import (
	"fmt"
)

type UniqueIndexViolation struct {
	Tablename string
	Err       error
	Indexname string
}

// QueryRowUnknown a db QueryRow has returned an error that hasn't been handled. Should result in a 500
type QueryRowUnknown struct {
	Tablename string
	Err       error
}

func (e *QueryRowUnknown) Error() string {
	return fmt.Sprintf("Unknown error querying %s: %v", e.Tablename, e.Err)
}
