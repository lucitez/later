package body

import (
	"github.com/google/uuid"
	"later/pkg/model"
)

// FriendRequestCreateBody ...
type FriendRequestCreateBody struct {
	SenderUserID    uuid.UUID
	RecipientUserID uuid.UUID
}

// ToFriendRequest ...
func (body *FriendRequestCreateBody) ToFriendRequest() (*model.FriendRequest, error) {
	return model.NewFriendRequest(
		body.SenderUserID,
		body.RecipientUserID)
}
