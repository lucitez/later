package response

import (
	"time"

	"github.com/google/uuid"
)

// WireFriendRequest
type WireFriendRequest struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
