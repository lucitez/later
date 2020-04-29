package body

import (
	"github.com/lucitez/later/api/src/pkg/model"

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
		ShareID:         shareID,
		ContentID:       body.Content.ID,
		RecipientUserID: body.RecipientUserID,
		SenderUserID:    body.SenderUserID,
	}
}
