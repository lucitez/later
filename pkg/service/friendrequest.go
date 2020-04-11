package service

import (
	"errors"
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/service/body"

	"github.com/google/uuid"
)

// FriendRequest ...
type FriendRequest struct {
	Repository    repository.FriendRequest
	FriendService Friend
	UserService   User
}

// NewFriendRequest for wire generation
func NewFriendRequest(
	repository repository.FriendRequest,
	friendService Friend,
	userService User,
) FriendRequest {
	return FriendRequest{
		repository,
		friendService,
		userService,
	}
}

// Create ...
func (manager *FriendRequest) Create(body body.FriendRequestCreateBody) (*model.FriendRequest, error) {
	friendRequest := body.ToFriendRequest()

	if err := manager.Repository.Insert(friendRequest); err != nil {
		return nil, err
	}

	return &friendRequest, nil
}

// Pending ...
func (manager *FriendRequest) Pending(userID uuid.UUID) []model.FriendRequest {
	return manager.Repository.PendingByUserID(userID)
}

// PendingBySenderAndRecipient ...
func (manager *FriendRequest) PendingBySenderAndRecipient(sentByUserID uuid.UUID, recipientUserID uuid.UUID) *model.FriendRequest {
	return manager.Repository.PendingByRequesterAndRequestee(sentByUserID, recipientUserID)
}

// Accept ...
func (manager *FriendRequest) Accept(id uuid.UUID) error {
	request := manager.Repository.ByID(id)

	if request == nil {
		return errors.New("Friend Request Not Found")
	}

	manager.Repository.Accept(id)
	return manager.FriendService.HandleAcceptedFriendRequest(*request)
}

// Decline ...
func (manager *FriendRequest) Decline(id uuid.UUID) {
	manager.Repository.Decline(id)
}
