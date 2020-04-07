package model

import (
	"database/sql"
	"log"
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserContent is the struct representing content that has been shared to a user. This is what shows up in their various feeds
type UserContent struct {
	ID           uuid.UUID
	ShareID      uuid.UUID
	ContentID    uuid.UUID
	ContentType  wrappers.NullString
	UserID       uuid.UUID
	SentByUserID uuid.UUID
	Tag          wrappers.NullString

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
	sentByUserID uuid.UUID,
) UserContent {

	id, _ := uuid.NewRandom()

	now := time.Now().UTC()

	userContent := UserContent{
		ID:           id,
		ShareID:      shareID,
		ContentID:    contentID,
		ContentType:  contentType,
		UserID:       userID,
		SentByUserID: sentByUserID,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return userContent
}

// ScanRows ...
func (userContent *UserContent) ScanRows(rows *sql.Rows) {
	err := rows.Scan(
		&userContent.ID,
		&userContent.ShareID,
		&userContent.ContentID,
		&userContent.ContentType,
		&userContent.UserID,
		&userContent.SentByUserID,
		&userContent.Tag,
		&userContent.CreatedAt,
		&userContent.UpdatedAt,
		&userContent.ArchivedAt,
		&userContent.DeletedAt,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (userContent *UserContent) ScanRow(row *sql.Row) *UserContent {
	err := row.Scan(
		&userContent.ID,
		&userContent.ShareID,
		&userContent.ContentID,
		&userContent.ContentType,
		&userContent.UserID,
		&userContent.SentByUserID,
		&userContent.Tag,
		&userContent.CreatedAt,
		&userContent.UpdatedAt,
		&userContent.ArchivedAt,
		&userContent.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return userContent
}
