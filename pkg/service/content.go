package service

import (
	"log"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/parse"
	"github.com/lucitez/later/pkg/repository"
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// Content ...
type Content struct {
	HostnameService Hostname
	Repository      repository.Content
}

// NewContent for wire generation
func NewContent(
	hostnameService Hostname,
	repo repository.Content,
) Content {
	return Content{
		hostnameService,
		repo,
	}
}

// GetContentPreview ...
func (service *Content) GetContentPreview(url string) (*parse.ContentMetadata, error) {
	return parse.ContentFromURL(url)
}

// CreateFromURL calls parse and creates content from parse results
func (service *Content) CreateFromURL(url string, userID uuid.UUID, contentType wrappers.NullString) (*model.Content, error) {
	contentMetadata, err := parse.ContentFromURL(url)

	if err != nil {
		return nil, err
	}

	content := contentMetadata.ToContent(userID, contentType)

	if err := service.Repository.Insert(content); err != nil {
		return nil, err
	}

	return &content, nil
}

// ByID returns Content by ID
func (service *Content) ByID(id uuid.UUID) (*model.Content, error) {
	return service.Repository.ByID(id)
}

// TODO is this safe from race conditions? probablly not.
// eventually set up a job to clean up share counts
func (service *Content) IncrementShareCount(id uuid.UUID, amount int) error {
	return service.Repository.IncrementShareCount(id, amount)
}

func (service *Content) TasteByUserID(userID uuid.UUID) int {
	taste, err := service.Repository.TasteByUserID(userID)

	if err != nil {
		log.Printf("[WARN] Error getting taste for user. Error: %s", err.Error())
	}

	return taste
}

// All returns all content within a limit
func (service *Content) All(limit int) ([]model.Content, error) {
	return service.Repository.All(limit)
}
