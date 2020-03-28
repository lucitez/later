package manager

import (
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
	"later.co/pkg/response"

	"github.com/google/uuid"
)

type FriendManager interface {
	All(userID uuid.UUID) ([]response.WireFriend, error)
	Search(userID uuid.UUID, query string) ([]response.WireFriend, error)
}

type FriendManagerImpl struct {
	UserManager UserManager
	Repository  repository.FriendRepository
}

// All ...
func (manager *FriendManagerImpl) All(userID uuid.UUID) ([]response.WireFriend, error) {
	friends, err := manager.Repository.ByUserID(userID)

	if err != nil {
		return nil, err
	}

	return manager.toWireFriends(friends), nil
}

// Search ...
func (manager *FriendManagerImpl) Search(userID uuid.UUID, query string) ([]response.WireFriend, error) {
	friends, err := manager.Repository.SearchByUserID(userID, query)

	if err != nil {
		return nil, err
	}

	return manager.toWireFriends(friends), nil
}

func (manager *FriendManagerImpl) toWireFriends(friends []entity.Friend) []response.WireFriend {
	wireFriends := []response.WireFriend{}

	for _, friend := range friends {
		friendUser, _ := manager.UserManager.ByID(friend.ID)
		if friendUser != nil {
			wireFriend := friend.ToWire(friendUser)
			wireFriends = append(wireFriends, wireFriend)
		}
	}

	return wireFriends
}
