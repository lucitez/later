package request

import (
	"github.com/google/uuid"
	"later.co/pkg/body"
)

// FriendRequestCreateRequestBody request body for sending a new friend request
type FriendRequestCreateRequestBody struct {
	SenderUserID    uuid.UUID `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientUserID uuid.UUID `form:"recipient_user_id" json:"recipient_user_ids" binding:"required"`
}

// ToFriendRequestCreateBody ...
func (requestBody *FriendRequestCreateRequestBody) ToFriendRequestCreateBody() body.FriendRequestCreateBody {
	return body.FriendRequestCreateBody{
		SenderUserID:    requestBody.SenderUserID,
		RecipientUserID: requestBody.RecipientUserID}
}

// FriendRequestAcceptRequestBody ...
type FriendRequestAcceptRequestBody struct {
	ID uuid.UUID `form:"id" json:"id" binding:"required"`
}

// FriendRequestDeclineRequestBody ...
type FriendRequestDeclineRequestBody struct {
	ID uuid.UUID `form:"id" json:"id" binding:"required"`
}
