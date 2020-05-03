package response

import (
	"time"

	"github.com/lucitez/later/pkg/util/wrappers"
)

type ContentPreview struct {
	URL         string              `json:"url"`
	Hostname    string              `json:"hostname"`
	Title       wrappers.NullString `json:"title"`
	Description wrappers.NullString `json:"description"`
	ImageURL    wrappers.NullString `json:"image_url"`
}

type PopularContent struct {
	URL         string
	Hostname    string
	Title       wrappers.NullString
	Description wrappers.NullString
	ImageURL    wrappers.NullString
	Shares      int
	CreatedAt   time.Time
}
