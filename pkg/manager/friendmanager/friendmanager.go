package friendmanager

import (
	"later.co/pkg/later/friend"
	"later.co/pkg/repository/userrepo"
	"later.co/pkg/response"

	"later.co/pkg/repository/friendrepo"

	"github.com/google/uuid"
)

func all(userID uuid.UUID) ([]response.WireFriend, error) {
	friends, err := friendrepo.ByUserID(userID)

	if err != nil {
		return nil, err
	}

	return friendsList{friends}.toWireFriends(), nil
}

func search(userID uuid.UUID, query string) ([]response.WireFriend, error) {
	friends, err := friendrepo.SearchByUserID(userID, query)

	if err != nil {
		return nil, err
	}

	return friendsList{friends}.toWireFriends(), nil
}

type friendsList struct {
	friends []friend.Friend
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
