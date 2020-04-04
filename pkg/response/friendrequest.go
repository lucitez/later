package response

import (
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type WireFriendRequest struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FirstName wrappers.NullString
	LastName  wrappers.NullString
	Username  wrappers.NullString
	CreatedAt time.Time
}
