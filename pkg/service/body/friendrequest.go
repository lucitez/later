package body

import (
	"later/pkg/model"

	"github.com/google/uuid"
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
