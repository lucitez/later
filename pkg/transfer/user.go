package transfer

import (
	"later/pkg/model"
	"later/pkg/response"
	"later/pkg/service"
	"later/pkg/util/wrappers"

	"github.com/google/uuid"
)

type User struct {
	FriendRequestService service.FriendRequest
	FriendService        service.Friend
}

func NewUser(
	friendRequestService service.FriendRequest,
	friendService service.Friend,
) User {
	return User{
		friendRequestService,
		friendService,
	}
}

func (transfer *User) WireUsersFrom(users []model.User) []response.WireUser {
	wireUsers := make([]response.WireUser, len(users))

	for i, user := range users {
		wireUsers[i] = transfer.WireUserFromUser(user)
	}

	return wireUsers
}

func (transfer *User) WireUserFromUser(user model.User) response.WireUser {
	return response.WireUser{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}
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
		Name:           user.Name,
		Username:       user.Username,
		PendingRequest: existingPendingRequest,
		CreatedAt:      user.CreatedAt,
	}
}

func (transfer *User) WireUserProfileFrom(requestUserID uuid.UUID, user model.User) response.WireUserProfile {
	wireUser := response.WireUserProfile{
		ID:           user.ID,
		Name:         user.Name,
		Username:     user.Username,
		FriendStatus: wrappers.NewNullString(nil),
	}

	if existingFriend := transfer.FriendService.ByUserIDAndFriendUserID(requestUserID, user.ID); existingFriend != nil {
		wireUser.FriendStatus = wrappers.NewNullStringFromString("friends")
		return wireUser
	}

	if pendingRequest := transfer.FriendRequestService.PendingBySenderAndRecipient(requestUserID, user.ID); pendingRequest != nil {
		wireUser.FriendStatus = wrappers.NewNullStringFromString("pending")
		wireUser.FriendRequestID = wrappers.NewNullUUIDFromUUID(pendingRequest.ID)
		return wireUser
	}

	return wireUser
}
