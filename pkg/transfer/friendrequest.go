package transfer

import (
	"later/pkg/model"
	"later/pkg/response"
	"later/pkg/service"
)

type FriendRequest struct {
	UserService service.User
}

func NewFriendRequest(userService service.User) FriendRequest {
	return FriendRequest{userService}
}

// WireFriendRequestsFrom tranfers DB model FriendRequest to DTO WireFriendRequest
func (transfer *FriendRequest) WireFriendRequestsFrom(friendRequests []model.FriendRequest) []response.WireFriendRequest {
	wireFriendRequests := make([]response.WireFriendRequest, len(friendRequests))

	for i, fr := range friendRequests {
		user := transfer.UserService.ByID(fr.SentByUserID)
		if user != nil {
			wireFriendRequests[i] = wireFriendRequest(fr, *user)
		}
	}

	return wireFriendRequests
}

func wireFriendRequest(fr model.FriendRequest, user model.User) response.WireFriendRequest {
	return response.WireFriendRequest{
		ID:        fr.ID,
		UserID:    fr.SentByUserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}
