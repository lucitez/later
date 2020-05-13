package body

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserContentCreateBody ...
type UserContentCreateBody struct {
	ShareID         uuid.UUID
	Content         model.Content
	RecipientUserID uuid.UUID
	Sender          model.User
}

// ToUserContent ...
func (body *UserContentCreateBody) ToUserContent() model.UserContent {
	return model.NewUserContent(
		body.ShareID,
		body.Content.ID,
		body.RecipientUserID,
		body.Sender.ID,
	)
}

// UserContentUpdateBody ...
type UserContentUpdateBody struct {
	ID  uuid.UUID
	Tag wrappers.NullString
}
