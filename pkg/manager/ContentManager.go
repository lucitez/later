package manager

import (
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
)

// ContentManager ...
type ContentManager struct {
	Repository repository.ContentRepository
}

// NewContentManager for wire generation
func NewContentManager(repo repository.ContentRepository) ContentManager {
	return ContentManager{repo}
}

// Create calls repository to create a new Content entry
func (manager *ContentManager) Create(content *entity.Content) (*entity.Content, error) {
	return manager.Repository.Insert(content)
}

// ByID returns Content by ID
func (manager *ContentManager) ByID(id uuid.UUID) (*entity.Content, error) {
	return manager.Repository.ByID(id)
}

// All returns all content within a limit
func (manager *ContentManager) All(limit int) ([]entity.Content, error) {
	return manager.Repository.All(limit)
}
