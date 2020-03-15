package body

import (
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
)

// ShareCreateBody ...
type ShareCreateBody struct {
	Content         entity.Content
	SenderUserID    uuid.UUID
	RecipientUserID uuid.UUID
}
