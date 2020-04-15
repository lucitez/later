package body

import (
	"later/pkg/model"
	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserContentCreateBody ...
type UserContentCreateBody struct {
	ShareID         uuid.UUID
	ContentID       uuid.UUID
	ContentType     wrappers.NullString
	RecipientUserID uuid.UUID
	SenderUserID    uuid.UUID
}

// ToUserContent ...
func (body *UserContentCreateBody) ToUserContent() model.UserContent {
	return model.NewUserContent(
		body.ShareID,
		body.ContentID,
		body.ContentType,
		body.RecipientUserID,
		body.SenderUserID,
	)
}

// UserContentUpdateBody ...
type UserContentUpdateBody struct {
	ID  uuid.UUID
	Tag wrappers.NullString
}
