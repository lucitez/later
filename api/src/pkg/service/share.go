package service

import (
	"github.com/lucitez/later/api/src/pkg/model"
	"github.com/lucitez/later/api/src/pkg/repository"
	"github.com/lucitez/later/api/src/pkg/service/body"
)

// Share ...
type Share struct {
	UserContent UserContent
	Message     Message
	Repository  repository.Share
}

// NewShare ...
func NewShare(
	repository repository.Share,
	userContent UserContent,
	message Message,
) Share {
	return Share{
		UserContent: userContent,
		Repository:  repository,
		Message:     message,
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

	go manager.UserContent.Create(body.ToUserContentCreateBody(share.ID))
	go manager.Message.CreateFromShare(share)

	return &share, nil
}
