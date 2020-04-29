package transfer

import (
	"later/pkg/model"
	"later/pkg/response"
	"later/pkg/service"
)

type Friend struct {
	UserService service.User
}

func NewFriend(userService service.User) Friend {
	return Friend{userService}
}

// WireFriendsFrom tranfers DB model Friend to DTO WireFriend
func (transfer *Friend) WireFriendsFrom(friendRequests []model.Friend) []response.WireFriend {
	wireFriends := make([]response.WireFriend, len(friendRequests))

	for i, friend := range friendRequests {
		user := transfer.UserService.ByID(friend.FriendUserID)
		if user != nil {
			wireFriends[i] = wireFriend(friend, *user)
		}
	}

	return wireFriends
}

func wireFriend(friend model.Friend, user model.User) response.WireFriend {
	return response.WireFriend{
		ID:          friend.ID,
		UserID:      user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}
}
