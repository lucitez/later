package response

import (
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type WireFriendRequest struct {
	ID        uuid.UUID           `json:"id"`
	UserID    uuid.UUID           `json:"user_id"`
	FirstName wrappers.NullString `json:"first_name"`
	LastName  wrappers.NullString `json:"last_name"`
	Username  wrappers.NullString `json:"username"`
	CreatedAt time.Time           `json:"created_at"`
}
