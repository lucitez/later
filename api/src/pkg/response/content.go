package response

import (
	"time"

	"github.com/lucitez/later/api/src/pkg/util/wrappers"

	"github.com/google/uuid"
)

// ContentHistory to be sent to client when displaing feed
type ContentHistory struct {
	ID          uuid.UUID           `json:"id"`
	Title       wrappers.NullString `json:"title"`
	Description wrappers.NullString `json:"description"`
	ImageURL    wrappers.NullString `json:"image_url"`
	ContentType wrappers.NullString `json:"content_type"`
	Shares      int                 `json:"taste_generated"`
	URL         string              `json:"url"`

	CreatedAt time.Time `json:"created_at"`
}

type ContentPreview struct {
	URL         string              `json:"url"`
	Title       wrappers.NullString `json:"title"`
	Description wrappers.NullString `json:"description"`
	ImageURL    wrappers.NullString `json:"image_url"`
	ContentType wrappers.NullString `json:"content_type"`
}
