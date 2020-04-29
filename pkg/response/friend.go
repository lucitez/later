package response

import (
	"time"

	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type WireFriend struct {
	ID          uuid.UUID           `json:"id"`
	UserID      uuid.UUID           `json:"user_id"`
	Name        string              `json:"name"`
	Username    string              `json:"username"`
	Email       wrappers.NullString `json:"email"`
	PhoneNumber string              `json:"phone_number"`
	CreatedAt   time.Time           `json:"created_at"`
}
