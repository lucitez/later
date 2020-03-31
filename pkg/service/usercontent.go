package service

import (
	"later/pkg/repository"
	"later/pkg/model"
	"later/pkg/response"

	"github.com/google/uuid"
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
func (manager *UserContentManager) Create(userContent *model.UserContent) (*model.UserContent, error) {
	return manager.Repository.Insert(userContent)
}

// ByID ...
func (manager *UserContentManager) ByID(id uuid.UUID) (*model.UserContent, error) {
	return manager.Repository.ByID(id)
}

// All ...
func (manager *UserContentManager) All(limit int) ([]model.UserContent, error) {
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
