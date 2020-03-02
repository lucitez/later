package body

import (
	"github.com/google/uuid"
	"later.co/pkg/later/content"
)

type ShareCreateBody struct {
	Content         content.Content
	SenderUserID    uuid.UUID
	RecipientUserID uuid.UUID
}
