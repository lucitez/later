package model

import (
	"database/sql"
	"log"
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// Content object
type Content struct {
	ID          uuid.UUID           `json:"id"`
	Title       wrappers.NullString `json:"title"`
	Description wrappers.NullString `json:"description"`
	ImageURL    wrappers.NullString `json:"image_url"`
	ContentType wrappers.NullString `json:"content_type"`
	URL         string              `json:"url"`
	Domain      string              `json:"domain"`
	Shares      int                 `json:"shares"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewContent constructor for Content
func NewContent(
	title wrappers.NullString,
	description wrappers.NullString,
	imageURL wrappers.NullString,
	contentType wrappers.NullString,
	url string,
	domain string,
) Content {
	id, _ := uuid.NewRandom()

	now := time.Now().UTC()

	content := Content{
		ID:          id,
		Title:       title,
		Description: description,
		ImageURL:    imageURL,
		ContentType: contentType,
		URL:         url,
		Domain:      domain,
		Shares:      0,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return content
}

// ScanRows ...
func (content *Content) ScanRows(rows *sql.Rows) {
	err := rows.Scan(
		&content.ID,
		&content.Title,
		&content.Description,
		&content.ImageURL,
		&content.ContentType,
		&content.URL,
		&content.Domain,
		&content.Shares,
		&content.CreatedAt,
		&content.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (content *Content) ScanRow(row *sql.Row) *Content {
	err := row.Scan(
		&content.ID,
		&content.Title,
		&content.Description,
		&content.ImageURL,
		&content.ContentType,
		&content.URL,
		&content.Domain,
		&content.Shares,
		&content.CreatedAt,
		&content.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return content
}
