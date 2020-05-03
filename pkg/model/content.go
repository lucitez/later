package model

import (
	"time"

	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// Content object
type Content struct {
	ID          uuid.UUID           `json:"id"`
	Title       wrappers.NullString `json:"title"`
	Description wrappers.NullString `json:"description"`
	ImageURL    wrappers.NullString `json:"image_url"`
	ContentType wrappers.NullString `json:"content_type"`
	URL         string              `json:"url"`
	Hostname    string              `json:"hostname"`
	Shares      int                 `json:"shares"`
	CreatedBy   uuid.UUID           `json:"created_by"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewContent constructor for Content
func NewContent(
	title wrappers.NullString,
	description wrappers.NullString,
	imageURL wrappers.NullString,
	contentType wrappers.NullString,
	url string,
	hostname string,
	createdBy uuid.UUID,
) Content {
	id, _ := uuid.NewRandom()

	now := time.Now().UTC()

	content := Content{
		ID:          id,
		Title:       title,
		Description: description,
		ImageURL:    imageURL,
		ContentType: contentType,
		URL:         url,
		Hostname:    hostname,
		Shares:      0,
		CreatedBy:   createdBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return content
}
