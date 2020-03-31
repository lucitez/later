package response

import (
	"time"

	"github.com/google/uuid"
	"later/pkg/util/wrappers"
)

type WireFriendRequest struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FirstName wrappers.NullString
	LastName  wrappers.NullString
	Username  wrappers.NullString
	CreatedAt time.Time
}
