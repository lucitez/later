package response

import (
	"time"

	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// WireUserContent to be sent to client when displaing feed
type WireUserContent struct {
	ID             uuid.UUID           `json:"id"`
	ContentID      uuid.UUID           `json:"content_id"`
	Title          wrappers.NullString `json:"title"`
	Description    wrappers.NullString `json:"description"`
	ImageURL       wrappers.NullString `json:"image_url"`
	ContentType    wrappers.NullString `json:"content_type"`
	Tag            wrappers.NullString `json:"tag"`
	Hostname       string              `json:"hostname"`
	Shares         int                 `json:"shares"`
	SentBy         uuid.UUID           `json:"sent_by"`
	SentByUsername string              `json:"sent_by_username"`
	SentByName     string              `json:"sent_by_name"`
	URL            string              `json:"url"`

	CreatedAt time.Time         `json:"created_at"`
	SavedAt   wrappers.NullTime `json:"saved_at"`
}
