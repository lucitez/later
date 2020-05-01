package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucitez/later/pkg/util/wrappers"
)

type UserMessage struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ChatID    uuid.UUID
	MessageID uuid.UUID

	CreatedAt time.Time
	ReadAt    wrappers.NullTime
	DeletedAt wrappers.NullTime
}

func NewUserMessage(
	chatID uuid.UUID,
	userID uuid.UUID,
	messageID uuid.UUID,
) UserMessage {
	ID, _ := uuid.NewRandom()

	now := time.Now().UTC()

	return UserMessage{
		ID:        ID,
		UserID:    userID,
		ChatID:    chatID,
		MessageID: messageID,
		CreatedAt: now,
	}
}
