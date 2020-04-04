package model

import (
	"database/sql"
	"log"
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
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
	recipientUserID uuid.UUID,
) FriendRequest {
	uuid, _ := uuid.NewRandom()

	now := time.Now().UTC()

	FriendRequest := FriendRequest{
		ID:              uuid,
		SentByUserID:    userID,
		RecipientUserID: recipientUserID,

		CreatedAt: now,
		UpdatedAt: now}

	return FriendRequest
}

// ScanRows ...
func (friendRequest *FriendRequest) ScanRows(rows *sql.Rows) {
	err := rows.Scan(
		&friendRequest.ID,
		&friendRequest.SentByUserID,
		&friendRequest.RecipientUserID,
		&friendRequest.CreatedAt,
		&friendRequest.UpdatedAt,
		&friendRequest.AcceptedAt,
		&friendRequest.DeclinedAt,
		&friendRequest.DeletedAt,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (friendRequest *FriendRequest) ScanRow(row *sql.Row) {
	err := row.Scan(
		&friendRequest.ID,
		&friendRequest.SentByUserID,
		&friendRequest.RecipientUserID,
		&friendRequest.CreatedAt,
		&friendRequest.UpdatedAt,
		&friendRequest.AcceptedAt,
		&friendRequest.DeclinedAt,
		&friendRequest.DeletedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
}
