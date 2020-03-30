package manager

import (
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
	"later.co/pkg/response"
)

// UserContentManager ...
type UserContentManager struct {
	Repository repository.UserContentRepository
}

// NewUserContentManager ...
func NewUserContentManager(repository repository.UserContentRepository) UserContentManager {
	return UserContentManager{repository}
}

// Create ...
func (manager *UserContentManager) Create(userContent *entity.UserContent) (*entity.UserContent, error) {
	return manager.Repository.Insert(userContent)
}

// ByID ...
func (manager *UserContentManager) ByID(id uuid.UUID) (*entity.UserContent, error) {
	return manager.Repository.ByID(id)
}

// All ...
func (manager *UserContentManager) All(limit int) ([]entity.UserContent, error) {
	return manager.Repository.All(limit)
}

// Feed ...
func (manager *UserContentManager) Feed(
	userID uuid.UUID,
	senderType *string,
	contentType *string,
	archived *bool) ([]response.WireUserContent, error) {

	return manager.Repository.Feed(
		userID,
		senderType,
		contentType,
		archived)
}
