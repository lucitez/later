package model

import (
	"later/pkg/util/wrappers"
	"time"

	"github.com/google/uuid"
)

// Chat struct for chats
type Chat struct {
	ID      uuid.UUID
	GroupID wrappers.NullUUID
	User1ID wrappers.NullUUID
	User2ID wrappers.NullUUID

	CreatedAt time.Time
	DeletedAt wrappers.NullTime
}

func NewGroupChat(
	groupID uuid.UUID,
) Chat {
	id, _ := uuid.NewRandom()

	now := time.Now().UTC()

	return Chat{
		ID:        id,
		GroupID:   wrappers.NewNullUUIDFromUUID(groupID),
		CreatedAt: now,
	}
}

func NewUserChat(
	user1ID uuid.UUID,
	user2ID uuid.UUID,
) Chat {
	id, _ := uuid.NewRandom()

	now := time.Now().UTC()

	return Chat{
		ID:        id,
		User1ID:   wrappers.NewNullUUIDFromUUID(user1ID),
		User2ID:   wrappers.NewNullUUIDFromUUID(user2ID),
		CreatedAt: now,
	}
}
