package response

import (
	"github.com/lucitez/later/pkg/util/wrappers"
)

type ContentPreview struct {
	URL         string              `json:"url"`
	Hostname    string              `json:"hostname"`
	Title       wrappers.NullString `json:"title"`
	Description wrappers.NullString `json:"description"`
	ImageURL    wrappers.NullString `json:"image_url"`
}
