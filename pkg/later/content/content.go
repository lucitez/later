package content

import (
	"time"

	"github.com/google/uuid"
	"later.co/pkg/util/wrappers"
)

// Content object
type Content struct {
	ID          uuid.UUID           `json:"id"`
	Title       string              `json:"title"`
	Description wrappers.NullString `json:"description"`
	ImageURL    wrappers.NullString `json:"image_url"`
	ContentType wrappers.NullString `json:"content_type"`
	URL         string              `json:"url"`
	Domain      string              `json:"domain"`
	Shares      int                 `json:"shares"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// New constructor for Content
func New(
	title string,
	description wrappers.NullString,
	imageURL wrappers.NullString,
	contentType wrappers.NullString,
	url string,
	domain string) (*Content, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	now := time.Now()

	content := Content{
		ID:          id,
		Title:       title,
		Description: description,
		ImageURL:    imageURL,
		ContentType: contentType,
		URL:         url,
		Domain:      domain,
		Shares:      0,
		CreatedAt:   now,
		UpdatedAt:   now}

	return &content, nil
}
