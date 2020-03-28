package manager

import (
	"later.co/pkg/body"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
)

type ShareManager interface {
	CreateMultiple(createBodies []body.ShareCreateBody) ([]entity.Share, error)
	Create(body body.ShareCreateBody) (*entity.Share, error)
}

type ShareManagerImpl struct {
	Repository         repository.ShareRepository
	UserContentManager UserContentManager
}

// CreateMultiple creates multiple shares from multiple bodies
func (manager *ShareManagerImpl) CreateMultiple(createBodies []body.ShareCreateBody) ([]entity.Share, error) {
	shares := []entity.Share{}

	for _, createBody := range createBodies {
		share, err := manager.Create(createBody)

		if err != nil {
			return nil, err
		}

		shares = append(shares, *share)
	}

	return shares, nil
}

// Create creates a share and usercontent
// Should probably do the notification stuff here
// TODO Two Goroutines:
// Update _body.Content.shares_ total by getting count(shares distinct on user_id with this content_id)
// Send Push notification if user has signed up <-- maybe move this to usercontent
func (manager *ShareManagerImpl) Create(body body.ShareCreateBody) (*entity.Share, error) {
	share, err := entity.NewShare(
		body.Content.ID,
		body.SenderUserID,
		body.RecipientUserID)

	if err != nil {
		return nil, err
	}

	share, err = manager.Repository.Insert(share)

	if err != nil {
		return nil, err
	}

	usercontent, err := entity.NewUserContent(
		share.ID,
		body.Content.ID,
		body.Content.ContentType,
		body.RecipientUserID,
		body.SenderUserID)

	if err != nil {
		return share, err
	}

	_, err = manager.UserContentManager.Create(usercontent)

	if err != nil {
		return nil, err
	}

	return share, nil
}
