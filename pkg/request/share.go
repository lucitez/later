package request

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/service/body"
	"github.com/lucitez/later/pkg/util/wrappers"

	"github.com/google/uuid"
)

// ShareCreateRequestBody Binding from json
// TODO experiment with null recipient user ids
type ShareCreateRequestBody struct {
	RecipientUserIDs []uuid.UUID         `form:"recipient_user_ids" json:"recipient_user_ids"`
	URL              string              `form:"url" json:"url" binding:"required"`
	ContentType      wrappers.NullString `json:"content_type"`
}

// ToShareCreateBodies converts this request body to a list of share create bodies
func (requestBody *ShareCreateRequestBody) ToShareCreateBodies(sender model.User, content model.Content) []body.ShareCreateBody {
	bodies := []body.ShareCreateBody{}

	for _, recipientUserID := range requestBody.RecipientUserIDs {
		createBody := body.ShareCreateBody{
			Content:         content,
			Sender:          sender,
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
func (requestBody *ShareForwardRequestBody) ToShareCreateBodies(sender model.User, content model.Content) []body.ShareCreateBody {
	bodies := []body.ShareCreateBody{}

	for _, recipientUserID := range requestBody.RecipientUserIDs {
		createBody := body.ShareCreateBody{
			Content:         content,
			Sender:          sender,
			RecipientUserID: recipientUserID,
		}

		bodies = append(bodies, createBody)
	}

	return bodies
}
