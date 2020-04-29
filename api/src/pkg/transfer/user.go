package transfer

import (
	"github.com/lucitez/later/api/src/pkg/model"
	"github.com/lucitez/later/api/src/pkg/response"
	"github.com/lucitez/later/api/src/pkg/service"
	"github.com/lucitez/later/api/src/pkg/util/wrappers"

	"github.com/google/uuid"
)

type User struct {
	FriendRequestService service.FriendRequest
	FriendService        service.Friend
	ContentService       service.Content
}

func NewUser(
	friendRequestService service.FriendRequest,
	friendService service.Friend,
	contentService service.Content,
) User {
	return User{
		friendRequestService,
		friendService,
		contentService,
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
	taste := transfer.ContentService.TasteByUserID(user.ID)

	return response.WireUser{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		Taste:       taste,
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
	taste := transfer.ContentService.TasteByUserID(user.ID)

	wireUser := response.WireUserProfile{
		ID:           user.ID,
		Name:         user.Name,
		Username:     user.Username,
		FriendStatus: wrappers.NewNullString(nil),
		Taste:        taste,
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
