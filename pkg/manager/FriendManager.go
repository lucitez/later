package manager

import (
	"later.co/pkg/later/entity"
	"later.co/pkg/repository"
	"later.co/pkg/response"

	"github.com/google/uuid"
)

// FriendManager ...
type FriendManager struct {
	UserManager UserManager
	Repository  repository.FriendRepository
}

// NewFriendManager for wire generation
func NewFriendManager(
	userManager UserManager,
	repository repository.FriendRepository) FriendManager {
	return FriendManager{
		UserManager: userManager,
		Repository:  repository}
}

// All ...
func (manager *FriendManager) All(userID uuid.UUID) ([]response.WireFriend, error) {
	friends, err := manager.Repository.ByUserID(userID)

	if err != nil {
		return nil, err
	}

	return manager.toWireFriends(friends), nil
}

// Search ...
func (manager *FriendManager) Search(userID uuid.UUID, query string) ([]response.WireFriend, error) {
	friends, err := manager.Repository.SearchByUserID(userID, query)

	if err != nil {
		return nil, err
	}

	return manager.toWireFriends(friends), nil
}

func (manager *FriendManager) toWireFriends(friends []entity.Friend) []response.WireFriend {
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
