package body

import (
	"github.com/google/uuid"
	"later/pkg/model"
)

// ShareCreateBody ...
type ShareCreateBody struct {
	Content         model.Content
	SenderUserID    uuid.UUID
	RecipientUserID uuid.UUID
}
