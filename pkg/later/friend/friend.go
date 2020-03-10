package friend

import (
	"time"

	"later.co/pkg/response"

	"github.com/google/uuid"
	"later.co/pkg/later/user"
	"later.co/pkg/util/wrappers"
)

// Friend object
type Friend struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	FriendUserID uuid.UUID `json:"friend_user_id"`

	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt wrappers.NullTime `json:"deleted_at"`
}

// New constructor for Friend
func New(
	userID uuid.UUID,
	friendUserID uuid.UUID) (*Friend, error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	now := time.Now()

	friend := Friend{
		ID:           uuid,
		UserID:       userID,
		FriendUserID: friendUserID,

		CreatedAt: now,
		UpdatedAt: now}

	return &friend, nil
}

func (friend *Friend) ToWire(friendUser *user.User) response.WireFriend {
	return response.WireFriend{
		ID:        friend.ID,
		UserID:    friendUser.ID,
		FirstName: friendUser.FirstName,
		LastName:  friendUser.LastName,
		Username:  friendUser.Username,
		CreatedAt: friend.CreatedAt}
}
