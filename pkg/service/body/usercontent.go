package body

import (
	"later/pkg/model"
	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type UserContentCreateBody struct {
	shareID         uuid.UUID
	contentID       uuid.UUID
	contentType     wrappers.NullString
	recipientUserID uuid.UUID
	senderUserID    uuid.UUID
}

func (body *UserContentCreateBody) ToUserContent() model.UserContent {
	return model.NewUserContent(
		body.shareID,
		body.contentID,
		body.contentType,
		body.recipientUserID,
		body.senderUserID,
	)
}
