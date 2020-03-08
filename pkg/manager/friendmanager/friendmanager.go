package friendmanager

import (
	"later.co/pkg/later/user"
	"later.co/pkg/repository/userrepo"

	"later.co/pkg/repository/friendrepo"

	"github.com/google/uuid"
)

func FriendsByUserID(userID uuid.UUID) ([]user.User, error) {
	users := []user.User{}

	friends, err := friendrepo.ByUserID(userID)

	if err != nil {
		return users, err
	}

	friendUserIds := make([]uuid.UUID, len(friends))

	for _, friend := range friends {
		friendUserIds = append(friendUserIds, friend.FriendUserID)
	}

	users, err = userrepo.ByIDs(friendUserIds)

	if err != nil {
		return users, err
	}

	return users, nil
}
