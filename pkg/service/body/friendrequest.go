package body

import (
	"github.com/lucitez/later/pkg/model"

	"github.com/google/uuid"
)

// FriendRequestCreateBody ...
type FriendRequestCreateBody struct {
	SenderUserID    uuid.UUID
	RecipientUserID uuid.UUID
}

// ToFriendRequest ...
func (body *FriendRequestCreateBody) ToFriendRequest() model.FriendRequest {
	return model.NewFriendRequest(
		body.SenderUserID,
		body.RecipientUserID,
	)
}
