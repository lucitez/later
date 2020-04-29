package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/lucitez/later/api/src/pkg/response"

	"github.com/lucitez/later/api/src/pkg/util/wrappers"

	"github.com/google/uuid"
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

// NewFriend constructor for Friend
func NewFriend(
	userID uuid.UUID,
	friendUserID uuid.UUID,
) Friend {
	uuid, _ := uuid.NewRandom()

	now := time.Now().UTC()

	friend := Friend{
		ID:           uuid,
		UserID:       userID,
		FriendUserID: friendUserID,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return friend
}

// ToWire transforms a Friend to a WireFriend
func (friend *Friend) ToWire(friendUser *User) response.WireFriend {
	return response.WireFriend{
		ID:        friend.ID,
		UserID:    friendUser.ID,
		Name:      friendUser.Name,
		Username:  friendUser.Username,
		CreatedAt: friend.CreatedAt,
	}
}

// ScanRows ...
func (friend *Friend) ScanRows(rows *sql.Rows) {
	err := rows.Scan(
		&friend.ID,
		&friend.UserID,
		&friend.FriendUserID,
		&friend.CreatedAt,
		&friend.UpdatedAt,
		&friend.DeletedAt,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (friend *Friend) ScanRow(row *sql.Row) *Friend {
	err := row.Scan(
		&friend.ID,
		&friend.UserID,
		&friend.FriendUserID,
		&friend.CreatedAt,
		&friend.UpdatedAt,
		&friend.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return friend
}
