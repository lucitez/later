package body

import (
	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserUpdate ...
type UserUpdate struct {
	ID          uuid.UUID
	FirstName   wrappers.NullString
	LastName    wrappers.NullString
	Email       wrappers.NullString
	PhoneNumber wrappers.NullString
}
