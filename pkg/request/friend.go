package request

import (
	"later/pkg/service/body"

	"github.com/google/uuid"
)

// FriendRequestCreateRequestBody request body for sending a new friend request
type FriendRequestCreateRequestBody struct {
	RecipientUserID uuid.UUID `form:"recipient_user_id" json:"recipient_user_id" binding:"required"`
}

// ToFriendRequestCreateBody ...
func (requestBody *FriendRequestCreateRequestBody) ToFriendRequestCreateBody(userID uuid.UUID) body.FriendRequestCreateBody {
	return body.FriendRequestCreateBody{
		SenderUserID:    userID,
		RecipientUserID: requestBody.RecipientUserID}
}

// FriendDeleteRequestBody request body for sending a new friend request
type FriendDeleteRequestBody struct {
	FriendUserID uuid.UUID `form:"friend_user_id" json:"friend_user_id" binding:"required"`
}

// FriendRequestAcceptRequestBody ...
type FriendRequestAcceptRequestBody struct {
	ID uuid.UUID `form:"id" json:"id" binding:"required"`
}

// FriendRequestDeclineRequestBody ...
type FriendRequestDeclineRequestBody struct {
	ID uuid.UUID `form:"id" json:"id" binding:"required"`
}

// FriendRequestDeleteRequestBody ...
type FriendRequestDeleteRequestBody struct {
	ID uuid.UUID `form:"id" json:"id" binding:"required"`
}
