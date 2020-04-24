package response

import (
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type WireFriend struct {
	ID          uuid.UUID           `json:"id"`
	UserID      uuid.UUID           `json:"user_id"`
	Name        wrappers.NullString `json:"name"`
	Username    wrappers.NullString `json:"username"`
	Email       wrappers.NullString `json:"email"`
	PhoneNumber string              `json:"phone_number"`
	CreatedAt   time.Time           `json:"created_at"`
}
