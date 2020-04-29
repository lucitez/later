package service

import (
	"later/pkg/model"
	"later/pkg/parse"
	"later/pkg/repository"
	"log"

	"github.com/google/uuid"
)

// Content ...
type Content struct {
	DomainService Domain
	Repository    repository.Content
}

// NewContent for wire generation
func NewContent(
	domainService Domain,
	repo repository.Content,
) Content {
	return Content{
		domainService,
		repo,
	}
}

// GetContentPreview ...
func (service *Content) GetContentPreview(url string) parse.ContentMetadata {
	contentMetadata := parse.ContentFromURL(url)

	return contentMetadata
}

// CreateFromURL calls parse and creates content from parse results
func (service *Content) CreateFromURL(url string, userID uuid.UUID) (*model.Content, error) {
	contentMetadata := parse.ContentFromURL(url)
	content := contentMetadata.ToContent(userID)

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
func (service *Content) IncrementShareCount(id uuid.UUID, amount int) error {
	return service.Repository.IncrementShareCount(id, amount)
}

func (service *Content) TasteByUserID(userID uuid.UUID) int {
	taste, err := service.Repository.TasteByUserID(userID)

	if err != nil {
		log.Printf("Error getting taste for user. Error: %s", err.Error())
	}

	return taste
}

// All returns all content within a limit
func (service *Content) All(limit int) []model.Content {
	return service.Repository.All(limit)
}