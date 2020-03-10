package response

import (
	"time"

	"github.com/google/uuid"
	"later.co/pkg/util/wrappers"
)

type WireFriend struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	FirstName   wrappers.NullString
	LastName    wrappers.NullString
	Username    wrappers.NullString
	Email       wrappers.NullString
	PhoneNumber string
	CreatedAt   time.Time
}
