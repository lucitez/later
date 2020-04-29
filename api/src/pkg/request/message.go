package request

import (
	"github.com/google/uuid"
)

type MessageSendRequestBody struct {
	ChatID  uuid.UUID `form:"chat_id" json:"chat_id" binding:"required"`
	Message string    `form:"message" json:"message" binding:"required"`
}
