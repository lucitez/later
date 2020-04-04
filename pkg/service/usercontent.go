package service

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/response"
	"later/pkg/service/body"

	"github.com/google/uuid"
)

// UserContent ...
type UserContent struct {
	Repository repository.UserContent
}

// NewUserContent ...
func NewUserContent(repository repository.UserContent) UserContent {
	return UserContent{repository}
}

// Create ...
func (manager *UserContent) Create(body body.UserContentCreateBody) (*model.UserContent, error) {
	userContent := body.ToUserContent()
	if err := manager.Repository.Insert(userContent); err != nil {
		return nil, err
	}
	return &userContent, nil
}

// ByID ...
func (manager *UserContent) ByID(id uuid.UUID) *model.UserContent {
	return manager.Repository.ByID(id)
}

// All ...
func (manager *UserContent) All(limit int) []model.UserContent {
	return manager.Repository.All(limit)
}

// Feed ...
func (manager *UserContent) Feed(
	userID uuid.UUID,
	senderType *string,
	contentType *string,
	archived *bool,
) ([]response.WireUserContent, error) {

	return manager.Repository.Feed(
		userID,
		senderType,
		contentType,
		archived)
}
