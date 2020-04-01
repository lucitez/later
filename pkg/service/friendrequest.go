package service

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/service/body"

	"github.com/google/uuid"
)

// FriendRequestManager ...
type FriendRequestManager struct {
	Repository repository.FriendRequest
}

// NewFriendRequestManager for wire generation
func NewFriendRequestManager(repository repository.FriendRequest) FriendRequestManager {
	return FriendRequestManager{repository}
}

// Create ...
func (manager *FriendRequestManager) Create(body body.FriendRequestCreateBody) (*model.FriendRequest, error) {
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
func (manager *FriendRequestManager) Pending(userID uuid.UUID) ([]model.FriendRequest, error) {
	requests, err := manager.Repository.PendingByUserID(userID)

	if err != nil {
		return nil, err
	}

	// wireFriendRequests := make([]response.WireFriendRequest, len(requests))

	return requests, err
}

// Accept ...
func (manager *FriendRequestManager) Accept(id uuid.UUID) error {
	return manager.Repository.Accept(id)
}

// Decline ...
func (manager *FriendRequestManager) Decline(id uuid.UUID) error {
	return manager.Repository.Decline(id)
}
