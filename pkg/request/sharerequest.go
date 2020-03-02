package request

import (
	"github.com/google/uuid"
	"later.co/pkg/body"
	"later.co/pkg/later/content"
)

// ShareCreateByUserIDRequestBody Binding from json
type ShareCreateByUserIDRequestBody struct {
	SenderUserID     uuid.UUID   `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientUserIDs []uuid.UUID `form:"recipient_user_ids" json:"recipient_user_ids" binding:"required"`
	URL              string      `form:"url" json:"url" binding:"required"`
}

// ToShareCreateBodies converts this request body to a list of share create bodies
func (requestBody *ShareCreateByUserIDRequestBody) ToShareCreateBodies(content *content.Content) []body.ShareCreateBody {
	bodies := []body.ShareCreateBody{}

	for _, recipientUserID := range requestBody.RecipientUserIDs {
		createBody := body.ShareCreateBody{
			Content:         *content,
			SenderUserID:    requestBody.SenderUserID,
			RecipientUserID: recipientUserID}

		bodies = append(bodies, createBody)
	}

	return bodies
}

// ShareCreateByPhoneNumberRequestBody Binding from json
type ShareCreateByPhoneNumberRequestBody struct {
	SenderUserID         uuid.UUID `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientPhoneNumber string    `form:"recipient_phone_number" json:"recipient_phone_number" binding:"required"`
	URL                  string    `form:"url" json:"url" binding:"required"`
}
