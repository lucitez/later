package response

import (
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type WireUser struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Username    string              `json:"username"`
	Email       wrappers.NullString `json:"email"`
	PhoneNumber string              `json:"phone_number"`
	Taste       int                 `json:"taste"`
	CreatedAt   time.Time           `json:"created_at"`
}

type WireAddFriendUser struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	PendingRequest bool      `json:"pending_request"`
	CreatedAt      time.Time `json:"created_at"`
}

type WireUserProfile struct {
	ID              uuid.UUID           `json:"id"`
	Name            string              `json:"name"`
	Username        string              `json:"username"`
	FriendStatus    wrappers.NullString `json:"friend_status"`
	FriendRequestID wrappers.NullUUID   `json:"friend_request_id"`
	Taste           int                 `json:"taste"`
}
