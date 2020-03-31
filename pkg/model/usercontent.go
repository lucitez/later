package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"later/pkg/util/wrappers"
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

// ScanRows ...
func (userContent *UserContent) ScanRows(rows *sql.Rows) error {
	err := rows.Scan(
		&userContent.ID,
		&userContent.ShareID,
		&userContent.ContentID,
		&userContent.ContentType,
		&userContent.UserID,
		&userContent.SentBy,
		&userContent.CreatedAt,
		&userContent.UpdatedAt,
		&userContent.ArchivedAt,
		&userContent.DeletedAt)

	return err
}

// ScanRow ...
func (userContent *UserContent) ScanRow(row *sql.Row) error {
	err := row.Scan(
		&userContent.ID,
		&userContent.ShareID,
		&userContent.ContentID,
		&userContent.ContentType,
		&userContent.UserID,
		&userContent.SentBy,
		&userContent.CreatedAt,
		&userContent.UpdatedAt,
		&userContent.ArchivedAt,
		&userContent.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}
