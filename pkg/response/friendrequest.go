package response

import (
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type WireFriendRequest struct {
	ID        uuid.UUID           `json:"id"`
	UserID    uuid.UUID           `json:"user_id"`
	Name      wrappers.NullString `json:"name"`
	Username  wrappers.NullString `json:"username"`
	CreatedAt time.Time           `json:"created_at"`
}
