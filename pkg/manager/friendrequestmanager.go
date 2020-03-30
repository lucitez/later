package manager

import (
	"github.com/google/uuid"
	"later.co/pkg/body"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
)

// FriendRequestManager ...
type FriendRequestManager struct {
	Repository repository.FriendRequestRepository
}

// NewFriendRequestManager for wire generation
func NewFriendRequestManager(repository repository.FriendRequestRepository) FriendRequestManager {
	return FriendRequestManager{repository}
}

// Create ...
func (manager *FriendRequestManager) Create(body body.FriendRequestCreateBody) (*entity.FriendRequest, error) {
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
func (manager *FriendRequestManager) Pending(userID uuid.UUID) ([]entity.FriendRequest, error) {
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
