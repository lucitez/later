package share

import (
	"time"

	"github.com/google/uuid"
	"later.co/pkg/util/wrappers"
)

// Share object
type Share struct {
	ID              uuid.UUID `json:"id"`
	ContentID       uuid.UUID `json:"content_id"`
	SentByUserID    uuid.UUID `json:"sent_by_user_id"`
	RecipientUserID uuid.UUID `json:"recipient_user_id"`

	CreatedAt time.Time         `json:"created_at"`
	OpenedAt  wrappers.NullTime `json:"opened_at"`
}

// New constructor for Share
func New(
	contentID uuid.UUID,
	sentByUserID uuid.UUID,
	recipientUserID uuid.UUID) (*Share, error) {

	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	now := time.Now()

	newShare := Share{
		ID:              uuid,
		ContentID:       contentID,
		SentByUserID:    sentByUserID,
		RecipientUserID: recipientUserID,

		CreatedAt: now}

	return &newShare, nil
}
