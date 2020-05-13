package body

import (
	"github.com/lucitez/later/pkg/model"

	"github.com/google/uuid"
)

// ShareCreateBody ...
type ShareCreateBody struct {
	Content         model.Content
	Sender          model.User
	RecipientUserID uuid.UUID
}

// ToUserContentCreateBody ...
func (body *ShareCreateBody) ToUserContentCreateBody(shareID uuid.UUID) UserContentCreateBody {
	return UserContentCreateBody{
		ShareID:         shareID,
		Content:         body.Content,
		RecipientUserID: body.RecipientUserID,
		Sender:          body.Sender,
	}
}
