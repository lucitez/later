package friendmanager

import (
	"later.co/pkg/later/entity"
	"later.co/pkg/response"

	"later.co/pkg/repository/friendrepo"

	"github.com/google/uuid"
)

// All ...
func All(userID uuid.UUID) ([]response.WireFriend, error) {
	friends, err := friendrepo.ByUserID(userID)

	if err != nil {
		return nil, err
	}

	return friendsList{friends}.toWireFriends(), nil
}

// Search ...
func Search(userID uuid.UUID, query string) ([]response.WireFriend, error) {
	friends, err := friendrepo.SearchByUserID(userID, query)

	if err != nil {
		return nil, err
	}

	return friendsList{friends}.toWireFriends(), nil
}

type friendsList struct {
	friends []entity.Friend
}

func (friendsList friendsList) toWireFriends() []response.WireFriend {
	wireFriends := []response.WireFriend{}

	for _, friend := range friendsList.friends {
		friendUser, _ := userrepo.ByID(friend.ID)
		if friendUser != nil {
			wireFriend := friend.ToWire(friendUser)
			wireFriends = append(wireFriends, wireFriend)
		}
	}

	return wireFriends
}
