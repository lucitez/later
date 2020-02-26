package request

import "github.com/google/uuid"

// ShareCreateRequestBody Binding from json
type ShareCreateRequestBody struct {
	SenderUserID    uuid.UUID `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientUserID uuid.UUID `form:"recipient_user_id" json:"recipient_user_id" binding:"required"`
	URL             string    `form:"url" json:"url" binding:"required"`
}
