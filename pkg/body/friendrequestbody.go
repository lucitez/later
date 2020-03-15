package body

import (
	"github.com/google/uuid"
	"later.co/pkg/later/entity"
)

// FriendRequestCreateBody ...
type FriendRequestCreateBody struct {
	SenderUserID    uuid.UUID
	RecipientUserID uuid.UUID
}

// ToFriendRequest ...
func (body *FriendRequestCreateBody) ToFriendRequest() (*entity.FriendRequest, error) {
	return entity.NewFriendRequest(
		body.SenderUserID,
		body.RecipientUserID)
}
