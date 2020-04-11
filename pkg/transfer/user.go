package transfer

import (
	"later/pkg/model"
	"later/pkg/response"
	"later/pkg/service"

	"github.com/google/uuid"
)

type User struct {
	FriendRequestService service.FriendRequest
}

func NewUser(friendRequestService service.FriendRequest) User {
	return User{friendRequestService}
}

// WireAddFriendUsersFrom tranfers DB model User to DTO WireAddFriendUser
func (transfer *User) WireAddFriendUsersFrom(userID uuid.UUID, users []model.User) []response.WireAddFriendUser {
	wireAddFriendUsers := make([]response.WireAddFriendUser, len(users))

	for i, user := range users {
		pendingFriendRequest := transfer.FriendRequestService.PendingBySenderAndRecipient(userID, user.ID)
		wireAddFriendUsers[i] = wireAddFriendUser(user, pendingFriendRequest != nil)

	}

	return wireAddFriendUsers
}

func wireAddFriendUser(user model.User, existingPendingRequest bool) response.WireAddFriendUser {
	return response.WireAddFriendUser{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		PendingRequest: existingPendingRequest,
		CreatedAt:      user.CreatedAt,
	}
}
