package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// Hostname object
type Hostname struct {
	ID          uuid.UUID `json:"id"`
	Hostname    string    `json:"hostname"`
	ContentType string    `json:"content_type"`

	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt wrappers.NullTime `json:"deleted_at"`
}

// NewHostname constructor for Hostname
func NewHostname(
	hostname string,
	contentType string,
) Hostname {
	uuid, _ := uuid.NewRandom()

	now := time.Now().UTC()

	newHostname := Hostname{
		ID:          uuid,
		Hostname:    hostname,
		ContentType: contentType,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return newHostname
}

// ScanRows ...
func (hostname *Hostname) ScanRows(rows *sql.Rows) {
	err := rows.Scan(
		&hostname.ID,
		&hostname.Hostname,
		&hostname.ContentType,
		&hostname.CreatedAt,
		&hostname.UpdatedAt,
		&hostname.DeletedAt)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (hostname *Hostname) ScanRow(row *sql.Row) *Hostname {
	err := row.Scan(
		&hostname.ID,
		&hostname.Hostname,
		&hostname.ContentType,
		&hostname.CreatedAt,
		&hostname.UpdatedAt,
		&hostname.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return hostname
}
