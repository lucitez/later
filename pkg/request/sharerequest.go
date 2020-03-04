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
	PhoneNumber      wrappers.NullString `form:"phone_number" json:"phone_number"`
	URL              string              `form:"url" json:"url" binding:"required"`
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

// ToShareCreateBodyByPhoneNumber converts this request body to a list of share create bodies
func (requestBody *ShareCreateRequestBody) ToShareCreateBodyByPhoneNumber(
	content *content.Content,
	recipientUserID uuid.UUID) body.ShareCreateBody {
	return body.ShareCreateBody{
		Content:         *content,
		SenderUserID:    requestBody.SenderUserID,
		RecipientUserID: recipientUserID}
}

// ShareForwardRequestBody Binding from json
type ShareForwardRequestBody struct {
	SenderUserID     uuid.UUID           `form:"sender_user_id" json:"sender_user_id" binding:"required"`
	RecipientUserIDs []uuid.UUID         `form:"recipient_user_ids" json:"recipient_user_ids"`
	PhoneNumber      wrappers.NullString `form:"phone_number" json:"phone_number"`
	ContentID        uuid.UUID           `form:"content_id" json:"content_id" binding:"required"`
}

// ToShareCreateBodies converts this request body to a list of share create bodies
func (requestBody *ShareForwardRequestBody) ToShareCreateBodies(content *content.Content) []body.ShareCreateBody {
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

// ToShareCreateBodyByPhoneNumber converts this request body to a list of share create bodies
func (requestBody *ShareForwardRequestBody) ToShareCreateBodyByPhoneNumber(
	content *content.Content,
	recipientUserID uuid.UUID) body.ShareCreateBody {
	return body.ShareCreateBody{
		Content:         *content,
		SenderUserID:    requestBody.SenderUserID,
		RecipientUserID: recipientUserID}
}
