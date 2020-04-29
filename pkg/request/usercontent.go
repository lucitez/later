package request

import (
	"github.com/lucitez/later/pkg/service/body"
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserContentSaveRequestBody ...
type UserContentSaveRequestBody struct {
	ID  uuid.UUID           `form:"id" json:"id" binding:"required"`
	Tag wrappers.NullString `form:"tag" json:"tag"`
}

// UserContentDeleteRequestBody ...
type UserContentDeleteRequestBody struct {
	ID uuid.UUID `form:"id" json:"id" binding:"required"`
}

// UserContentUpdateRequestBody ...
type UserContentUpdateRequestBody struct {
	ID  uuid.UUID           `form:"id" json:"id" binding:"required"`
	Tag wrappers.NullString `form:"tag" json:"tag"`
}

// ToUserContentUpdateBody ...
func (requestBody *UserContentUpdateRequestBody) ToUserContentUpdateBody() body.UserContentUpdateBody {
	return body.UserContentUpdateBody{
		ID:  requestBody.ID,
		Tag: requestBody.Tag,
	}
}
