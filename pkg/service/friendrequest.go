package service

import (
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
	friendRequest, err := body.ToFriendRequest()

	if err != nil {
		return nil, err
	}

	friendRequest, err = manager.Repository.Insert(friendRequest)

	if err != nil {
		return nil, err
	}

	return friendRequest, nil
}

// Pending ...
func (manager *FriendRequest) Pending(userID uuid.UUID) ([]model.FriendRequest, error) {
	requests, err := manager.Repository.PendingByUserID(userID)

	if err != nil {
		return nil, err
	}

	return requests, err
}

// Accept ...
func (manager *FriendRequest) Accept(id uuid.UUID) error {
	var err error

	request, err := manager.Repository.ByID(id)

	if err != nil {
		return err
	}

	if request != nil {
		err = manager.Repository.Accept(id)
		err = manager.FriendService.HandleAcceptedFriendRequest(*request)
	}

	return err
}

// Decline ...
func (manager *FriendRequest) Decline(id uuid.UUID) error {
	return manager.Repository.Decline(id)
}
