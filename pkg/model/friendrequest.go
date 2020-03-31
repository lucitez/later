package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"later/pkg/util/wrappers"
)

// FriendRequest object
type FriendRequest struct {
	ID              uuid.UUID `json:"id"`
	SentByUserID    uuid.UUID `json:"sent_by_user_id"`
	RecipientUserID uuid.UUID `json:"recipient_user_id"`

	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	AcceptedAt wrappers.NullTime `json:"accepted_at"`
	DeclinedAt wrappers.NullTime `json:"declined_at"`
	DeletedAt  wrappers.NullTime `json:"deleted_at"`
}

// NewFriendRequest constructor for FriendRequest
func NewFriendRequest(
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

// ScanRows ...
func (friendRequest *FriendRequest) ScanRows(rows *sql.Rows) error {
	err := rows.Scan(
		&friendRequest.ID,
		&friendRequest.SentByUserID,
		&friendRequest.RecipientUserID,
		&friendRequest.CreatedAt,
		&friendRequest.UpdatedAt,
		&friendRequest.AcceptedAt,
		&friendRequest.DeclinedAt,
		&friendRequest.DeletedAt)

	return err
}

// ScanRow ...
func (friendRequest *FriendRequest) ScanRow(row *sql.Row) error {
	err := row.Scan(
		&friendRequest.ID,
		&friendRequest.SentByUserID,
		&friendRequest.RecipientUserID,
		&friendRequest.CreatedAt,
		&friendRequest.UpdatedAt,
		&friendRequest.AcceptedAt,
		&friendRequest.DeclinedAt,
		&friendRequest.DeletedAt)

	return err
}
