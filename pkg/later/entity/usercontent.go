package entity

import (
	"time"

	"github.com/google/uuid"
	"later.co/pkg/util/wrappers"
)

// UserContent is the struct representing content that has been shared to a user. This is what shows up in their various feeds
type UserContent struct {
	ID          uuid.UUID
	ShareID     uuid.UUID
	ContentID   uuid.UUID
	ContentType wrappers.NullString
	UserID      uuid.UUID
	SentBy      uuid.UUID

	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt wrappers.NullTime
	DeletedAt  wrappers.NullTime
}

// NewUserContent constructor for UserContent
func NewUserContent(
	shareID uuid.UUID,
	contentID uuid.UUID,
	contentType wrappers.NullString,
	userID uuid.UUID,
	sentBy uuid.UUID) (*UserContent, error) {

	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	now := time.Now()

	userContent := UserContent{
		ID:          id,
		ShareID:     shareID,
		ContentID:   contentID,
		ContentType: contentType,
		UserID:      userID,
		SentBy:      sentBy,

		CreatedAt: now,
		UpdatedAt: now}

	return &userContent, nil
}
