package response

import (
	"later/pkg/util/wrappers"
	"time"

	"github.com/google/uuid"
)

type WireMessage struct {
	ID      uuid.UUID
	ChatID  uuid.UUID
	Message wrappers.NullString
	Content WireMessageContent
	SentBy  uuid.UUID
	SentAt  time.Time
}

type WireMessageContent struct {
	ID          uuid.UUID           `json:"id"`
	Title       wrappers.NullString `json:"title"`
	Description wrappers.NullString `json:"description"`
	ImageURL    wrappers.NullString `json:"image_url"`
	URL         string              `json:"url"`
}
