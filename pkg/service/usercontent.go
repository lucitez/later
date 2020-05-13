package service

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"github.com/lucitez/later/pkg/service/body"
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// UserContent ...
type UserContent struct {
	repo                repository.UserContent
	notificationService Notification
}

// NewUserContent ...
func NewUserContent(
	repository repository.UserContent,
	notificationService Notification,
) UserContent {
	return UserContent{
		repository,
		notificationService,
	}
}

// Create ...
func (uc *UserContent) Create(body body.UserContentCreateBody) (*model.UserContent, error) {
	userContent := body.ToUserContent()
	if err := uc.repo.Insert(userContent); err != nil {
		return nil, err
	}

	go uc.sendContentSharedNotification(body.Content, body.Sender, body.RecipientUserID)
	return &userContent, nil
}

func (uc *UserContent) sendContentSharedNotification(content model.Content, from model.User, to uuid.UUID) {
	notificationMessage := PushMessage{
		To:    to,
		Title: from.Name + " shared content with you",
		Body:  content.Title.String + " - " + content.URL,
	}

	uc.notificationService.SendMessage(notificationMessage)
}

// ByID ...
func (uc *UserContent) ByID(id uuid.UUID) *model.UserContent {
	return uc.repo.ByID(id)
}

// All ...
func (uc *UserContent) All(limit int) []model.UserContent {
	return uc.repo.All(limit)
}

// Filter ...
func (uc *UserContent) Filter(
	userID uuid.UUID,
	tag *string,
	contentType *string,
	saved bool,
	search *string,
	limit int,
) ([]model.UserContent, error) {

	return uc.repo.Filter(
		userID,
		tag,
		contentType,
		saved,
		search,
		limit,
	)
}

// FilterTags ...
func (uc *UserContent) FilterTags(
	userID uuid.UUID,
	search *string,
) ([]string, error) {

	return uc.repo.FilterTags(
		userID,
		search,
	)
}

// ByTag ...
func (uc *UserContent) ByTag(
	userID uuid.UUID,
	tag string,
) ([]model.UserContent, error) {

	return uc.repo.ByTag(
		userID,
		tag,
	)
}

// Save a piece of user content, providing an optional tag
func (uc *UserContent) Save(
	id uuid.UUID,
	tag wrappers.NullString,
) error {
	return uc.repo.Save(id, tag)
}

// Delete a post
func (uc *UserContent) Delete(id uuid.UUID) error {
	return uc.repo.Delete(id)
}

// Update user content
func (uc *UserContent) Update(body body.UserContentUpdateBody) {
	uc.repo.Update(body)
}
