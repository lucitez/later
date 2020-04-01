package model

import (
	"database/sql"
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// Content object
type Content struct {
	ID          uuid.UUID           `json:"id"`
	Title       string              `json:"title"`
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
	title string,
	description wrappers.NullString,
	imageURL wrappers.NullString,
	contentType wrappers.NullString,
	url string,
	domain string) (*Content, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

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
		UpdatedAt:   now}

	return &content, nil
}

// ScanRows ...
func (content *Content) ScanRows(rows *sql.Rows) error {
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
		&content.UpdatedAt)

	return err
}

// ScanRow ...
func (content *Content) ScanRow(row *sql.Row) error {
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
		&content.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}
