package request

import (
	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserContentArchiveRequestBody ...
type UserContentArchiveRequestBody struct {
	ID  uuid.UUID           `form:"id" json:"id" binding:"required"`
	Tag wrappers.NullString `form:"tag" json:"tag"`
}

// UserContentDeleteRequestBody ...
type UserContentDeleteRequestBody struct {
	ID uuid.UUID `form:"id" json:"id" binding:"required"`
}
