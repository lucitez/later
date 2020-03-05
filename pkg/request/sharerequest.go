package request

import (
	"github.com/google/uuid"
	"later.co/pkg/body"
	"later.co/pkg/later/content"
	"later.co/pkg/util/wrappers"
)

// ShareCreateRequestBody Binding from json
// TODO experiment with null recipient user ids
type ShareCreateRequestBody struct {
	SenderUserID     uuid.UUID           `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientUserIDs []uuid.UUID         `form:"recipient_user_ids" json:"recipient_user_ids"`
	URL              wrappers.NullString `form:"url" json:"url"`               // If new content
	ContentID        wrappers.NullUUID   `form:"content_id" json:"content_id"` // If forwarding
}

// ToShareCreateBodiesByUserIds converts this request body to a list of share create bodies
func (requestBody *ShareCreateRequestBody) ToShareCreateBodiesByUserIds(content *content.Content) []body.ShareCreateBody {
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
	SenderUserID uuid.UUID           `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	PhoneNumber  string              `form:"phone_number" json:"phone_number" binding:"required"`
	URL          wrappers.NullString `form:"url" json:"url" binding:"required"` // If new content
	ContentID    wrappers.NullUUID   `form:"content_id" json:"content_id"`      // If forwarding
}
