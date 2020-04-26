package service

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/service/body"
)

// Share ...
type Share struct {
	Content     Content
	UserContent UserContent
	Repository  repository.Share
}

// NewShare ...
func NewShare(
	repository repository.Share,
	userContent UserContent,
) Share {
	return Share{
		UserContent: userContent,
		Repository:  repository,
	}
}

// Create creates a share and usercontent
// TODO Send Push notification if user has signed up <-- maybe move this to usercontent
func (manager *Share) Create(body body.ShareCreateBody) (*model.Share, error) {
	share := model.NewShare(
		body.Content.ID,
		body.SenderUserID,
		body.RecipientUserID,
	)

	if err := manager.Repository.Insert(share); err != nil {
		return nil, err
	}

	if _, err := manager.UserContent.Create(body.ToUserContentCreateBody(share.ID)); err != nil {
		return nil, err
	}

	return &share, nil
}
