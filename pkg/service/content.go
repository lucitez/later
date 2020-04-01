package service

import (
	"github.com/google/uuid"
	"later/pkg/model"
	"later/pkg/repository"
)

// ContentManager ...
type ContentManager struct {
	Repository repository.Content
}

// NewContentManager for wire generation
func NewContentManager(repo repository.Content) ContentManager {
	return ContentManager{repo}
}

// Create calls repository to create a new Content entry
func (manager *ContentManager) Create(content *model.Content) (*model.Content, error) {
	return manager.Repository.Insert(content)
}

// ByID returns Content by ID
func (manager *ContentManager) ByID(id uuid.UUID) (*model.Content, error) {
	return manager.Repository.ByID(id)
}

// All returns all content within a limit
func (manager *ContentManager) All(limit int) ([]model.Content, error) {
	return manager.Repository.All(limit)
}
