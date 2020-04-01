package model

import (
	"database/sql"
	"time"

	"later/pkg/response"

	"github.com/google/uuid"
	"later/pkg/util/wrappers"
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

// ToWire transforms a Friend to a WireFriend
func (friend *Friend) ToWire(friendUser *User) response.WireFriend {
	return response.WireFriend{
		ID:        friend.ID,
		UserID:    friendUser.ID,
		FirstName: friendUser.FirstName,
		LastName:  friendUser.LastName,
		Username:  friendUser.Username,
		CreatedAt: friend.CreatedAt}
}

// ScanRows ...
func (friend *Friend) ScanRows(rows *sql.Rows) error {
	err := rows.Scan(
		&friend.ID,
		&friend.UserID,
		&friend.FriendUserID,
		&friend.CreatedAt,
		&friend.UpdatedAt,
		&friend.DeletedAt)

	return err
}

// ScanRow ...
func (friend *Friend) ScanRow(row *sql.Row) error {
	err := row.Scan(
		&friend.ID,
		&friend.UserID,
		&friend.FriendUserID,
		&friend.CreatedAt,
		&friend.UpdatedAt,
		&friend.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}