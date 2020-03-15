package entity

import (
	"time"

	"github.com/google/uuid"
	"later.co/pkg/util/wrappers"
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

	now := time.Now()

	newDomain := Domain{
		ID:          uuid,
		Domain:      domain,
		ContentType: contentType,

		CreatedAt: now,
		UpdatedAt: now}

	return &newDomain, nil
}
