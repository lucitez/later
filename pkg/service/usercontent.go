package service

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/service/body"
	"later/pkg/util/wrappers"

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
func (service *UserContent) Create(body body.UserContentCreateBody) (*model.UserContent, error) {
	userContent := body.ToUserContent()
	if err := service.Repository.Insert(userContent); err != nil {
		return nil, err
	}
	return &userContent, nil
}

// ByID ...
func (service *UserContent) ByID(id uuid.UUID) *model.UserContent {
	return service.Repository.ByID(id)
}

// All ...
func (service *UserContent) All(limit int) []model.UserContent {
	return service.Repository.All(limit)
}

// Filter ...
func (service *UserContent) Filter(
	userID uuid.UUID,
	tag *string,
	contentType *string,
	saved bool,
	search *string,
	limit int,
) []model.UserContent {

	return service.Repository.Filter(
		userID,
		tag,
		contentType,
		saved,
		search,
		limit,
	)
}

// FilterTags ...
func (service *UserContent) FilterTags(
	userID uuid.UUID,
	search *string,
) ([]string, error) {

	return service.Repository.FilterTags(
		userID,
		search,
	)
}

// ByTag ...
func (service *UserContent) ByTag(
	userID uuid.UUID,
	tag string,
) ([]model.UserContent, error) {

	return service.Repository.ByTag(
		userID,
		tag,
	)
}

// Save a piece of user content, providing an optional tag
func (service *UserContent) Save(
	id uuid.UUID,
	tag wrappers.NullString,
) error {
	return service.Repository.Save(id, tag)
}

// Delete a post
func (service *UserContent) Delete(id uuid.UUID) error {
	return service.Repository.Delete(id)
}

// Update user content
func (service *UserContent) Update(body body.UserContentUpdateBody) {
	service.Repository.Update(body)
}
