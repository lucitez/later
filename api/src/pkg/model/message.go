package model

import (
	"github.com/lucitez/later/api/src/pkg/util/wrappers"
	"time"

	"github.com/google/uuid"
)

// Message struct for message
type Message struct {
	ID        uuid.UUID
	ChatID    uuid.UUID
	SentBy    uuid.UUID
	Message   wrappers.NullString
	ContentID wrappers.NullUUID

	CreatedAt time.Time
	DeletedAt wrappers.NullTime
}

func NewMessage(
	chatID uuid.UUID,
	sentBy uuid.UUID,
	message string,
) Message {
	ID, _ := uuid.NewRandom()

	now := time.Now().UTC()

	return Message{
		ID:      ID,
		ChatID:  chatID,
		SentBy:  sentBy,
		Message: wrappers.NewNullStringFromString(message),

		CreatedAt: now,
	}
}

func NewMessageFromContent(
	chatID uuid.UUID,
	sentBy uuid.UUID,
	contentID uuid.UUID,
) Message {
	ID, _ := uuid.NewRandom()

	now := time.Now().UTC()

	return Message{
		ID:        ID,
		ChatID:    chatID,
		SentBy:    sentBy,
		ContentID: wrappers.NewNullUUIDFromUUID(contentID),

		CreatedAt: now,
	}
}
