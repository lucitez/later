package friendrequest

import (
	"time"

	"github.com/google/uuid"
	"later.co/pkg/util/wrappers"
)

// FriendRequest object
type FriendRequest struct {
	ID              uuid.UUID `json:"id"`
	SentByUserID    uuid.UUID `json:"user_id"`
	RecipientUserID uuid.UUID `json:"FriendRequest_user_id"`

	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	AcceptedAt wrappers.NullTime `json:"accepted_at"`
	DeclinedAt wrappers.NullTime `json:"declined_at"`
	DeletedAt  wrappers.NullTime `json:"deleted_at"`
}

// New constructor for FriendRequest
func New(
	userID uuid.UUID,
	recipientUserID uuid.UUID) (*FriendRequest, error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	now := time.Now()

	FriendRequest := FriendRequest{
		ID:              uuid,
		SentByUserID:    userID,
		RecipientUserID: recipientUserID,

		CreatedAt: now,
		UpdatedAt: now}

	return &FriendRequest, nil
}
