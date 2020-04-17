package response

import (
	"time"

	"later/pkg/util/wrappers"

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
	Domain         string              `json:"domain"`
	SentBy         uuid.UUID           `json:"sent_by"`
	SentByUsername wrappers.NullString `json:"sent_by_username"`
	URL            string              `json:"url"`

	CreatedAt  time.Time         `json:"created_at"`
	SavedAt wrappers.NullTime `json:"saved_at"`
}
