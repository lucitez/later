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
	repository repository.Friend) Friend {
	return Friend{
		User:       userManager,
		Repository: repository}
}

// HandleAcceptedFriendRequest creates two new friend entries. One for the requester, one for the requestee.
func (manager *Friend) HandleAcceptedFriendRequest(request model.FriendRequest) error {
	requester, err := model.NewFriend(
		request.SentByUserID,
		request.RecipientUserID,
	)
	requestee, err := model.NewFriend(
		request.RecipientUserID,
		request.SentByUserID,
	)
	_, err = manager.Repository.Insert(requester)
	_, err = manager.Repository.Insert(requestee)

	return err
}

// All ...
func (manager *Friend) All(userID uuid.UUID) []model.Friend {
	return manager.Repository.ByUserID(userID)
}

// Search ...
func (manager *Friend) Search(userID uuid.UUID, query string) []model.Friend {
	return manager.Repository.SearchByUserID(userID, query)
}
