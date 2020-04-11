package service

import (
	"later/pkg/model"
	"later/pkg/repository"

	"github.com/google/uuid"
)

// Friend ...
type Friend struct {
	User       User
	Repository repository.Friend
}

// NewFriend for wire generation
func NewFriend(
	userManager User,
	repository repository.Friend,
) Friend {
	return Friend{
		User:       userManager,
		Repository: repository,
	}
}

// HandleAcceptedFriendRequest creates two new friend entries. One for the requester, one for the requestee.
func (manager *Friend) HandleAcceptedFriendRequest(request model.FriendRequest) error {
	requester := model.NewFriend(
		request.SentByUserID,
		request.RecipientUserID,
	)
	requestee := model.NewFriend(
		request.RecipientUserID,
		request.SentByUserID,
	)
	if err := manager.Repository.Insert(requester); err != nil {
		return err
	}
	if err := manager.Repository.Insert(requestee); err != nil {
		return err
	}

	return nil
}

// All ...
func (manager *Friend) ForUser(
	userID uuid.UUID,
	search *string,
	limit int,
	offset int,
) []model.Friend {
	return manager.Repository.ForUser(
		userID,
		search,
		limit,
		offset,
	)
}
