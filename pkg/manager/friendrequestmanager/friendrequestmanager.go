package friendrequestmanager

import (
	"github.com/google/uuid"
	"later.co/pkg/body"
	"later.co/pkg/later/entity"
	"later.co/pkg/repository/friendrequestrepo"
)

// Create ...
func Create(body body.FriendRequestCreateBody) (*entity.FriendRequest, error) {
	friendRequest, err := body.ToFriendRequest()

	if err != nil {
		return nil, err
	}

	friendRequest, err = friendrequestrepo.Insert(friendRequest)

	if err != nil {
		return nil, err
	}

	return friendRequest, nil
}

// Pending ...
func Pending(userID uuid.UUID) ([]entity.FriendRequest, error) {
	return friendrequestrepo.PendingByUserID(userID)
}

// Accept ...
func Accept(id uuid.UUID) error {
	return friendrequestrepo.Accept(id)
}

// Decline ...
func Decline(id uuid.UUID) error {
	return friendrequestrepo.Decline(id)
}
