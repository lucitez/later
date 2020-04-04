package request

import (
	"later/pkg/model"
	"later/pkg/service/body"

	"github.com/google/uuid"
)

// ShareCreateRequestBody Binding from json
// TODO experiment with null recipient user ids
type ShareCreateRequestBody struct {
	SenderUserID     uuid.UUID   `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientUserIDs []uuid.UUID `form:"recipient_user_ids" json:"recipient_user_ids"`
	URL              string      `form:"url" json:"url" binding:"required"`
}

// ShareForwardRequestBody Binding from json
// TODO experiment with null recipient user ids
type ShareForwardRequestBody struct {
	SenderUserID     uuid.UUID   `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientUserIDs []uuid.UUID `form:"recipient_user_ids" json:"recipient_user_ids"`
	ContentID        uuid.UUID   `form:"content_id" json:"content_id" binding:"required"`
}

// ToShareCreateBodiesByUserIds converts this request body to a list of share create bodies
func (requestBody *ShareCreateRequestBody) ToShareCreateBodiesByUserIds(content *model.Content) []body.ShareCreateBody {
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
// TODO experiment with null recipient user ids
type ShareCreateByPhoneNumberRequestBody struct {
	SenderUserID uuid.UUID `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	PhoneNumber  string    `form:"phone_number" json:"phone_number" binding:"required"`
	URL          string    `form:"url" json:"url" binding:"required"`
}
