package body

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserContentCreateBody ...
type UserContentCreateBody struct {
	ShareID         uuid.UUID
	ContentID       uuid.UUID
	RecipientUserID uuid.UUID
	SenderUserID    uuid.UUID
}

// ToUserContent ...
func (body *UserContentCreateBody) ToUserContent() model.UserContent {
	return model.NewUserContent(
		body.ShareID,
		body.ContentID,
		body.RecipientUserID,
		body.SenderUserID,
	)
}

// UserContentUpdateBody ...
type UserContentUpdateBody struct {
	ID  uuid.UUID
	Tag wrappers.NullString
}
