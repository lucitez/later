package transfer

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/response"
	"github.com/lucitez/later/pkg/service"
)

type Message struct {
	ContentService service.Content
}

func NewMessage(cs service.Content) Message {
	return Message{cs}
}

func (t *Message) WireMessagesFrom(messages []model.Message) []response.WireMessage {
	wireMessages := make([]response.WireMessage, len(messages))

	for i, message := range messages {
		wireMessages[i] = t.WireMessageFromMessage(message)
	}

	return wireMessages
}

func (t *Message) WireMessageFromMessage(message model.Message) response.WireMessage {
	wireMessage := response.WireMessage{
		ID:      message.ID,
		ChatID:  message.ChatID,
		SentBy:  message.SentBy,
		SentAt:  message.CreatedAt,
		Message: message.Message,
	}

	if message.ContentID.Valid {
		if content, _ := t.ContentService.ByID(message.ContentID.ID); content != nil {
			wireContent := response.WireMessageContent{
				ID:          content.ID,
				Title:       content.Title,
				Description: content.Description,
				ImageURL:    content.ImageURL,
				URL:         content.URL,
			}
			wireMessage.Content = wireContent
		}
	}

	return wireMessage
}
