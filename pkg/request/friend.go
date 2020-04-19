package request

import (
	"later/pkg/service/body"

	"github.com/google/uuid"
)

// FriendRequestCreateRequestBody request body for sending a new friend request
type FriendRequestCreateRequestBody struct {
	SenderUserID    uuid.UUID `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientUserID uuid.UUID `form:"recipient_user_id" json:"recipient_user_id" binding:"required"`
}

// ToFriendRequestCreateBody ...
func (requestBody *FriendRequestCreateRequestBody) ToFriendRequestCreateBody() body.FriendRequestCreateBody {
	return body.FriendRequestCreateBody{
		SenderUserID:    requestBody.SenderUserID,
		RecipientUserID: requestBody.RecipientUserID}
}

// FriendDeleteRequestBody request body for sending a new friend request
type FriendDeleteRequestBody struct {
	UserID       uuid.UUID `form:"user_id" json:"user_id" binding:"required"`
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
