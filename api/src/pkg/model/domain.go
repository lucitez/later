package model

import (
	"database/sql"
	"log"
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
	contentType string,
) Domain {
	uuid, _ := uuid.NewRandom()

	now := time.Now().UTC()

	newDomain := Domain{
		ID:          uuid,
		Domain:      domain,
		ContentType: contentType,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return newDomain
}

// ScanRows ...
func (domain *Domain) ScanRows(rows *sql.Rows) {
	err := rows.Scan(
		&domain.ID,
		&domain.Domain,
		&domain.ContentType,
		&domain.CreatedAt,
		&domain.UpdatedAt,
		&domain.DeletedAt)

	if err != nil {
		log.Fatal(err)
	}
}

// ScanRow ...
func (domain *Domain) ScanRow(row *sql.Row) *Domain {
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
		panic(err)
	}

	return domain
}
