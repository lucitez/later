package manager

import (
	"github.com/google/uuid"
	"later.co/pkg/body"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
)

type FriendRequestManager interface {
	Create(body body.FriendRequestCreateBody) (*entity.FriendRequest, error)
	Pending(userID uuid.UUID) ([]entity.FriendRequest, error)
	Accept(id uuid.UUID) error
	Decline(id uuid.UUID) error
}

type FriendRequestManagerImpl struct {
	Repository repository.FriendRequestRepository
}

// Create ...
func (manager *FriendRequestManagerImpl) Create(body body.FriendRequestCreateBody) (*entity.FriendRequest, error) {
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
func (manager *FriendRequestManagerImpl) Pending(userID uuid.UUID) ([]entity.FriendRequest, error) {
	requests, err := manager.Repository.PendingByUserID(userID)

	if err != nil {
		return nil, err
	}

	// wireFriendRequests := make([]response.WireFriendRequest, len(requests))

	return requests, err
}

// Accept ...
func (manager *FriendRequestManagerImpl) Accept(id uuid.UUID) error {
	return manager.Repository.Accept(id)
}

// Decline ...
func (manager *FriendRequestManagerImpl) Decline(id uuid.UUID) error {
	return manager.Repository.Decline(id)
}
