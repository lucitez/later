package request

import (
	"later/pkg/model"
	"later/pkg/service/body"

	"github.com/google/uuid"
)

// ShareCreateRequestBody Binding from json
// TODO experiment with null recipient user ids
type ShareCreateRequestBody struct {
	RecipientUserIDs []uuid.UUID `form:"recipient_user_ids" json:"recipient_user_ids"`
	URL              string      `form:"url" json:"url" binding:"required"`
}

// ToShareCreateBodies converts this request body to a list of share create bodies
func (requestBody *ShareCreateRequestBody) ToShareCreateBodies(senderUserID uuid.UUID, content model.Content) []body.ShareCreateBody {
	bodies := []body.ShareCreateBody{}

	for _, recipientUserID := range requestBody.RecipientUserIDs {
		createBody := body.ShareCreateBody{
			Content:         content,
			SenderUserID:    senderUserID,
			RecipientUserID: recipientUserID,
		}

		bodies = append(bodies, createBody)
	}

	return bodies
}

// ShareForwardRequestBody Binding from json
// TODO experiment with null recipient user ids
type ShareForwardRequestBody struct {
	RecipientUserIDs []uuid.UUID `form:"recipient_user_ids" json:"recipient_user_ids"`
	ContentID        uuid.UUID   `form:"content_id" json:"content_id" binding:"required"`
}

// ToShareCreateBodies this request body to a list of share create bodies
func (requestBody *ShareForwardRequestBody) ToShareCreateBodies(senderUserID uuid.UUID, content model.Content) []body.ShareCreateBody {
	bodies := []body.ShareCreateBody{}

	for _, recipientUserID := range requestBody.RecipientUserIDs {
		createBody := body.ShareCreateBody{
			Content:         content,
			SenderUserID:    senderUserID,
			RecipientUserID: recipientUserID,
		}

		bodies = append(bodies, createBody)
	}

	return bodies
}

// ShareCreateByPhoneNumberRequestBody Binding from json
type ShareCreateByPhoneNumberRequestBody struct {
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"required"`
	URL         string `form:"url" json:"url" binding:"required"`
}
