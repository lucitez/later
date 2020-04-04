package model

import (
	"database/sql"
	"log"
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
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

// NewShare constructor for Share
func NewShare(
	contentID uuid.UUID,
	sentByUserID uuid.UUID,
	recipientUserID uuid.UUID,
) Share {

	uuid, _ := uuid.NewRandom()

	now := time.Now().UTC()

	newShare := Share{
		ID:              uuid,
		ContentID:       contentID,
		SentByUserID:    sentByUserID,
		RecipientUserID: recipientUserID,

		CreatedAt: now,
	}

	return newShare
}

// ScanRows ...
func (share *Share) ScanRows(rows *sql.Rows) {
	err := rows.Scan(
		&share.ID,
		&share.ContentID,
		&share.SentByUserID,
		&share.RecipientUserID,
		&share.CreatedAt,
		&share.OpenedAt)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (share *Share) ScanRow(row *sql.Row) *Share {
	err := row.Scan(
		&share.ID,
		&share.ContentID,
		&share.SentByUserID,
		&share.RecipientUserID,
		&share.CreatedAt,
		&share.OpenedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return share
}
