package model

import (
	"database/sql"
	"time"

	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// Domain object
type Domain struct {
	ID          uuid.UUID `json:"id"`
	Domain      string    `json:"domain"`
	ContentType string    `json:"content_type"`

	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt wrappers.NullTime `json:"deleted_at"`
}

// NewDomain constructor for Domain
func NewDomain(
	domain string,
	contentType string) (*Domain, error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	newDomain := Domain{
		ID:          uuid,
		Domain:      domain,
		ContentType: contentType,

		CreatedAt: now,
		UpdatedAt: now}

	return &newDomain, nil
}

// ScanRows ...
func (domain *Domain) ScanRows(rows *sql.Rows) error {
	err := rows.Scan(
		&domain.ID,
		&domain.Domain,
		&domain.ContentType,
		&domain.CreatedAt,
		&domain.UpdatedAt,
		&domain.DeletedAt)

	return err
}

// ScanRow ...
func (domain *Domain) ScanRow(row *sql.Row) error {
	err := row.Scan(
		&domain.ID,
		&domain.Domain,
		&domain.ContentType,
		&domain.CreatedAt,
		&domain.UpdatedAt,
		&domain.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return err
}
