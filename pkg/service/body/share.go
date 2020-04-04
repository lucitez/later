package body

import (
	"later/pkg/model"

	"github.com/google/uuid"
)

// ShareCreateBody ...
type ShareCreateBody struct {
	Content         model.Content
	SenderUserID    uuid.UUID
	RecipientUserID uuid.UUID
}

// ToUserContentCreateBody ...
func (body *ShareCreateBody) ToUserContentCreateBody(shareID uuid.UUID) UserContentCreateBody {
	return UserContentCreateBody{
		shareID,
		body.Content.ID,
		body.Content.ContentType,
		body.RecipientUserID,
		body.SenderUserID,
	}
}
