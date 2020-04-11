package response

import (
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type WireAddFriendUser struct {
	ID             uuid.UUID           `json:"id"`
	FirstName      wrappers.NullString `json:"first_name"`
	LastName       wrappers.NullString `json:"last_name"`
	Username       wrappers.NullString `json:"username"`
	PendingRequest bool                `json:"pending_request"`
	CreatedAt      time.Time           `json:"created_at"`
}
